package controller

import (
	"strconv"

	"github.com/Youngkingman/Kchat/kchat/global"
	"github.com/Youngkingman/Kchat/kchat/internal/model"
	"github.com/Youngkingman/Kchat/kchat/internal/service"
	"github.com/Youngkingman/Kchat/kchat/pkg/app"
	"github.com/Youngkingman/Kchat/kchat/pkg/errcode"
	"github.com/gin-gonic/gin"
)

// 前端需要的数据
type respUser struct {
	Uid    int    `json:"uid"`
	Name   string `json:"name"`
	Avatar string `json:"img"`
}

// The users who are allow to this chat room
func ChatroomUserList(c *gin.Context) {
	resp := app.NewResponse(c)
	ridStr := c.Query("rid")
	rid, err := strconv.Atoi(ridStr)
	if err != nil {
		global.Logger.Errorf(c, "get chatroom info fail with error: %v", err)
		resp.ToErrorResponse(errcode.TransStringFail.WithDetails(err.Error()))
		return
	}
	usersList, err := model.GetAllUserFromChatRoom(c, rid)
	if err != nil {
		global.Logger.Errorf(c, "get users from chat %v fail with error: %v", rid, err)
		resp.ToErrorResponse(errcode.ErrorGetRoomUserFail.WithDetails(err.Error()))
		return
	}
	// 转换成前端想要的
	respUsers := make([]*respUser, 0)
	for _, v := range usersList {
		tmp := &respUser{
			Uid:    v.UID,
			Name:   v.Name,
			Avatar: v.ImageURL,
		}
		respUsers = append(respUsers, tmp)
	}
	resp.ToResponse(respUsers)
}

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
