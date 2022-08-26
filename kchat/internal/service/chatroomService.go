package service

import (
	"context"

	"github.com/Youngkingman/Kchat/kchat/global"
	"github.com/Youngkingman/Kchat/kchat/internal/model"
	"github.com/Youngkingman/Kchat/kchat/pkg/errcode"
)

type ChatRoom struct {
	ChatRoom  *model.ChatRoom
	BroadCast *broadcaster
}

// 根据房间号索引对应的房间
// 1.先校验用户是否注册过房间以及房间是否存在
// 2.启动时直接加载所有房间，项目小问题不大，项目大emmmm分布式分配聊天房间
// 缺少用户注册接口，测试先自己强开一个
// 跨域问题需要解决,不然整不动vue-cli

func LoadAllChatRoom() error {
	ctx := context.Background()
	rooms, err := model.GetAllChatRoom(ctx)
	if err != nil {
		global.Logger.Panicf(ctx, "load rooms failed with error %v", err)
		return err
	}
	for _, v := range rooms {
		cr := &ChatRoom{
			ChatRoom:  v,
			BroadCast: NewBroadCast(v.RoomID),
		}
		// 全局加载的时候启动对应广播器的协程进行监听，只要服务器运行协程就不退出
		go cr.BroadCast.Start()
		ChatRoomMap.Set(v.RoomID, cr)
	}
	return nil
}

// 刷新加载特定房间，用于新用户加入时的刷新
func LoadChatRoom(rid int) (err error) {
	ctx := context.Background()
	room, err := model.GetChatRoomByRoomId(ctx, rid)
	if err != nil {
		global.Logger.Panicf(ctx, "load room %d failed with error %v", rid, err)
		return err
	}
	// ChatRoomMap.Set(rid,room)
	v, ok := ChatRoomMap.Get(rid)
	if ok {
		v.ChatRoom = room
	} else {
		// 理论上走不到这里，但以后说不定
		// cr := &ChatRoom{
		// 	ChatRoom:  room,
		// 	BroadCast: NewBroadCast(rid),
		// }
		// // 全局加载的时候启动对应广播器的协程进行监听，只要服务器运行协程就不退出
		// go cr.BroadCast.Start()
		// ChatRoomMap.Set(rid, cr)
		err = errcode.ErrorGetChatRoomInfoFail
	}
	return err
}
