package controller

import (
	"strconv"

	"github.com/Youngkingman/Kchat/kchat/global"
	"github.com/Youngkingman/Kchat/kchat/internal/model"
	"github.com/Youngkingman/Kchat/kchat/internal/service"
	"github.com/Youngkingman/Kchat/kchat/pkg/app"
	"github.com/Youngkingman/Kchat/kchat/pkg/errcode"
	"github.com/gin-gonic/gin"
	"nhooyr.io/websocket"
)

func ChatroomWebsocket(c *gin.Context) {
	// 查看jwt实现，该接口必须鉴权再进入,接口以查询形式给予参数
	u := c.MustGet("user").(*model.User)
	ridStr := c.Query("rid")
	rid, err := strconv.Atoi(ridStr)
	resp := app.NewResponse(c)
	if err != nil {
		global.Logger.Errorf(c, "get chatroom info fail with: %v, no rid parsed", err)
		resp.ToErrorResponse(errcode.TransStringFail.WithDetails(err.Error()))
		return
	}
	// 查看对应房间是否存在,此时需要控制并发
	// 但是我们后面调用方法的时候不需要考虑，为什么呢？
	// 因为广播器的方法都是对channel的一层良好封装，它们并发安全
	chatroom, ok := service.ChatRoomMap.Get(rid)
	if !ok {
		global.Logger.Debugf(c, "can't find chat room with rid %v", rid)
		resp.ToErrorResponse(errcode.ErrorGetChatRoomInfoFail.WithDetails(err.Error()))
		return
	}
	// 检查用户是否有进入房间的权利
	v, ok := chatroom.ChatRoom.Users[u.UID]
	if !ok {
		global.Logger.Debugf(c, "user %d no right to enter room %d", v, rid)
		resp.ToErrorResponse(errcode.ErrorNoRightToAccessRoom.WithDetails(err.Error()))
		return
	}

	// 对跨域进行处理和一般的跨域不同，需要设置 `InsecureSkipVerify` 选项
	// 一个小坑，由于http writer会被websocket库劫持，所以所有校验必须在建立websocket连接之前完成，否则报500的服务器错误
	conn, err := websocket.Accept(c.Writer, c.Request, &websocket.AcceptOptions{InsecureSkipVerify: true})
	if err != nil {
		global.Logger.Errorf(c, "websocket connect fail with err %v", err)
		resp.ToErrorResponse(errcode.WebsocketConnectFail)
		return
	}

	broadcaster := chatroom.BroadCast
	// 1.根据入参构造该房间新的用户实例
	chatter := service.NewChatter(u, conn, rid)
	// 2.开启``服务器``给``用户``发送消息的goroutine
	go chatter.SendMessage(c)
	// 3.服务器给用户发送欢迎信息
	chatter.MessageChannel <- service.MsgSrv.NewWelcomeMessage(chatter.Chatter)
	// 4. 用户加入该房间广播器，并告知所有人，我觉得可以暂时去掉
	broadcaster.ChatterEntering(chatter)
	// global.Logger.Debugf(c, "chatter:%v enter the room %v", chatter.Chatter.Name, rid)
	// msg := service.MsgSrv.NewChatterEnterMessage(chatter.Chatter)
	// broadcaster.Broadcast(c, msg)
	// 5.用户开始接受消息，该函数为循环阻塞函数
	err = chatter.ReceiveMessage(c, broadcaster)

	// 6. 循环阻塞函数退出，用户离开
	broadcaster.ChatterLeaving(chatter)
	msg := service.MsgSrv.NewChatterLeaveMessage(chatter.Chatter)
	broadcaster.Broadcast(c, msg)
	global.Logger.Debugf(c, "chatter:%v leave the room %v", chatter.Chatter.Name, rid)

	if err == nil {
		conn.Close(websocket.StatusNormalClosure, "")
	} else {
		global.Logger.Errorf(c, "read from client error: %v", err)
		conn.Close(websocket.StatusInternalError, "Read from client error")
	}
}
