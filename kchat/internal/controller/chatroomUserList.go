package controller

import (
	"strconv"

	"github.com/Youngkingman/Kchat/kchat/global"
	"github.com/Youngkingman/Kchat/kchat/internal/service"
	"github.com/Youngkingman/Kchat/kchat/pkg/app"
	"github.com/Youngkingman/Kchat/kchat/pkg/errcode"
	"github.com/gin-gonic/gin"
)

// The user who are allow to this chat room
// func ChatroomUserList(c *gin.Context) {
// }

func ChatRoomChatterList(c *gin.Context) {
	resp := app.NewResponse(c)
	ridStr := c.Query("rid")
	rid, err := strconv.Atoi(ridStr)
	// 查询列表带上房间号
	if err != nil {
		global.Logger.Errorf(c, "get chatroom info fail with error: %v", err)
		resp.ToErrorResponse(errcode.TransStringFail.WithDetails(err.Error()))
		return
	}
	chatroom, ok := service.ChatRoomMap.Get(rid)
	if !ok {
		global.Logger.Debugf(c, "can't find chat room with rid %v", rid)
		resp.ToErrorResponse(errcode.ErrorGetChatRoomInfoFail.WithDetails(err.Error()))
		return
	}
	chatterList := chatroom.BroadCast.GetChatterList()
	resp.ToResponse(chatterList)
}
