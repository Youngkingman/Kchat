package controller

import (
	"github.com/Youngkingman/Kchat/kchat/global"
	"github.com/Youngkingman/Kchat/kchat/internal/model"
	"github.com/Youngkingman/Kchat/kchat/pkg/app"
	"github.com/Youngkingman/Kchat/kchat/pkg/errcode"
	"github.com/Youngkingman/Kchat/kchat/pkg/psw"
	"github.com/gin-gonic/gin"
)

type signupReq struct {
	Name           string `json:"name" binding:"required,gte=1,lte=30"`
	Email          string `json:"email" binding:"required,email"`
	Password       string `json:"password" binding:"required,gte=6,lte=30"`
	RepeatPassword string `json:"repeatPassword" binding:"required,gte=6,lte=30"`
}

func Signup(c *gin.Context) {
	var req signupReq
	resp := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &req)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		resp.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	if req.Password != req.RepeatPassword {
		resp.ToErrorResponse(errcode.ErrorRepeatPswInconsist)
		return
	}

	hashPw, err := psw.HashPassword(req.Password)
	if err != nil {
		global.Logger.Errorf(c, "fail to hash password %v", err)
		resp.ToErrorResponse(errcode.ErrorHashPasswordFail.WithDetails(err.Error()))
		return
	}
	// 默认用户头像
	u := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashPw,
		ImageURL: global.AppSetting.DefaultAvatar,
	}
	err = model.AddUser(c, u)
	if err != nil {
		global.Logger.Errorf(c, "fail to sign up with errs %v", err)
		resp.ToErrorResponse(errcode.ErrorDuplicateUserWithEmail.WithDetails(err.Error()))
		return
	}
	// here we need to return a token, finish later
	// 生成token
	token, err := app.GenerateToken(u)
	if err != nil {
		resp.ToErrorResponse(errcode.ErrorTokenGenerateFail)
		return
	}
	// 需要回传一下用户信息用于初次登录时的前端vuex状态缓存
	u, err = model.GetUserByEmail(c, u.Email)
	if err != nil {
		resp.ToErrorResponse(errcode.ErrorGetUserInfoFail)
		return
	}
	// 将token加入redis中
	err = model.SetToken(u.Email, token)
	// 回传token
	resp.ToResponse(gin.H{"token": token, "user": u})
}
