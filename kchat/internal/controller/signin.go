package controller

import (
	"github.com/Youngkingman/Kchat/kchat/global"
	"github.com/Youngkingman/Kchat/kchat/internal/model"
	"github.com/Youngkingman/Kchat/kchat/pkg/app"
	"github.com/Youngkingman/Kchat/kchat/pkg/errcode"
	"github.com/Youngkingman/Kchat/kchat/pkg/psw"
	"github.com/gin-gonic/gin"
)

type signinReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=6,lte=30"`
}

func Signin(c *gin.Context) {
	var req signinReq
	resp := app.NewResponse(c)
	// 获取登录信息
	valid, errs := app.BindAndValid(c, &req)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		resp.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	// 查找密码并校验
	u, err := model.GetUserByEmail(c, req.Email)
	if err != nil {
		global.Logger.Errorf(c, "fail to sign in with errs %v", err)
		resp.ToErrorResponse(errcode.ErrorDuplicateUserWithEmail.WithDetails(err.Error()))
		return
	}
	storePsw := u.Password
	ok, err := psw.ComparePasswords(storePsw, req.Password)

	if err != nil {
		global.Logger.Errorf(c, "fail to sign in with errs %v", err)
		resp.ToErrorResponse(errcode.ErrorHashPasswordFail.WithDetails(err.Error()))
		return
	}
	if !ok {
		global.Logger.Errorf(c, "wrong password, fail to sign in with errs %v", err)
		resp.ToErrorResponse(errcode.ErrorInvalidPassword)
		return
	}
	// 可以检查数据库中是否有token并提示风险，这里单纯每次都更新
	// 生成token
	token, err := app.GenerateToken(u)
	if err != nil {
		global.Logger.Errorf(c, "fail to generate token with errs %v", err)
		resp.ToErrorResponse(errcode.ErrorTokenGenerateFail)
		return
	}
	// 将token加入redis中
	err = model.SetToken(u.Email, token)
	// 回传tokenh和用户数据
	u.Password = ""
	if err != nil {
		global.Logger.Errorf(c, "fail to set token with errs %v", err)
		resp.ToErrorResponse(errcode.ErrorTokenGenerateFail)
		return
	}

	resp.ToResponse(gin.H{"token": token, "user": u})
}
