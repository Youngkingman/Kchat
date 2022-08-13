package controller

import (
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
	}
	resp.ToResponse(nil)
}

func AddUserSToChatRoomTest(c *gin.Context) {
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
		global.Logger.Errorf(c, "add chat room fail with error: %v", err)
		resp.ToErrorResponse(errcode.ErrorAddChatRoomFail.WithDetails(err.Error()))
	}
}
