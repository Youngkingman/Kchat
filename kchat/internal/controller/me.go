package controller

import (
	"log"

	"github.com/Youngkingman/Kchat/kchat/global"
	"github.com/Youngkingman/Kchat/kchat/internal/model"
	"github.com/Youngkingman/Kchat/kchat/pkg/app"
	"github.com/Youngkingman/Kchat/kchat/pkg/errcode"
	"github.com/gin-gonic/gin"
)

// 一个没啥用的测试接口，正常一个前端传过来一个user的混杂uidStr后端解码之后
func Me(c *gin.Context) {
	// A *model.User will eventually be added to context in middleware
	user, exists := c.Get("user")

	// This shouldn't happen, as our middleware ought to throw an error.
	// This is an extra safety measure
	req := app.NewResponse(c)
	if !exists {
		// global.Logger.Debugf(c, "Unable to extract user from request context for unknown reason: %v\n", c)
		// // err := apperrors.NewInternal()
		// req.ToErrorResponse(errcode.ServerError.WithDetails("This shouldn't happen, as our middleware ought to throw an error."))

		// return
	}

	uid := user.(*model.User).UID
	if !exists {
		uid = 1
	}

	// use the Request Context
	// ctx := c.Request.Context()
	u, err := model.GetUser(c.Request.Context(), uid)

	if err != nil {
		log.Printf("Unable to find user: %v\n%v", uid, err)
		global.Logger.Debugf(c, "Unable to find user: %v\n%v", uid, err)
		req.ToErrorResponse(errcode.ErrorGetUserInfoFail)

		return
	}
	req.ToResponse(u)
}
