package controller

import (
	"log"
	"strconv"

	"github.com/Youngkingman/Kchat/kchat/global"
	"github.com/Youngkingman/Kchat/kchat/internal/model"
	"github.com/Youngkingman/Kchat/kchat/pkg/app"
	"github.com/Youngkingman/Kchat/kchat/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type AddChatRoomReq struct {
	Name string `json:"name"`
	Uids []int  `json:"uids"`
}

type AddUserSToChatRoomReq struct {
	Rid  int   `json:"rid"`
	Uids []int `json:"uids"`
}

func AddChatRoomTest(c *gin.Context) {
	var req AddChatRoomReq
	resp := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &req)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		resp.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	err := model.AddChatRoom(c, req.Name, req.Uids)
	if err != nil {
		global.Logger.Errorf(c, "add chat room fail with error: %v", err)
		resp.ToErrorResponse(errcode.ErrorAddChatRoomFail.WithDetails(err.Error()))
		return
	}
	resp.ToResponse(nil)
}

func AddUserSToChatRoom(c *gin.Context) {
	var req AddUserSToChatRoomReq
	resp := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &req)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		resp.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	err := model.AddUserSToChatRoom(c, req.Rid, req.Uids)
	if err != nil {
		global.Logger.Errorf(c, "add chatroom fail with error: %v", err)
		resp.ToErrorResponse(errcode.ErrorAddChatRoomFail.WithDetails(err.Error()))
		return
	}
	resp.ToResponse(nil)
}

func GetChatRoomByRoomIdTest(c *gin.Context) {
	ridStr := c.Query("rid")
	resp := app.NewResponse(c)
	rid, err := strconv.Atoi(ridStr)
	if err != nil {
		global.Logger.Errorf(c, "get chatroom info fail with error: %v", err)
		resp.ToErrorResponse(errcode.TransStringFail.WithDetails(err.Error()))
		return
	}
	chatRoom, err := model.GetChatRoomByRoomId(c, rid)
	if err != nil {
		global.Logger.Errorf(c, "get chatroom info fail with error: %v", err)
		resp.ToErrorResponse(errcode.ErrorGetChatRoomInfoFail.WithDetails(err.Error()))
		return
	}
	log.Println(*chatRoom)
	//resp.ToResponse(*chatRoom)
}

func GetAllChatRoomTest(c *gin.Context) {
	resp := app.NewResponse(c)
	chatRooms, err := model.GetAllChatRoom(c)
	if err != nil {
		global.Logger.Errorf(c, "get chatroom info fail with error: %v", err)
		resp.ToErrorResponse(errcode.ErrorGetChatRoomInfoFail.WithDetails(err.Error()))
		return
	}
	ret := make([]model.ChatRoom, 0)
	for _, v := range chatRooms {
		ret = append(ret, *v)
	}
	log.Println(chatRooms)
	//resp.ToResponse(ret)
}
