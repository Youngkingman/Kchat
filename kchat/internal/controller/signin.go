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

	valid, errs := app.BindAndValid(c, &req)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		resp.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	storePsw, err := model.GetUserPasswordByEmail(c, req.Email)
	psw.ComparePasswords(storePsw, req.Password)

	if err != nil {
		global.Logger.Errorf(c, "fail to sign up with errs %v", err)
		resp.ToErrorResponse(errcode.ErrorSignUpFail.WithDetails(err.Error()))
		return
	}
	// here we need to return a token, finish later
	resp.ToResponse("sign up success")
}
