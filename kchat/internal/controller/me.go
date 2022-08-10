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

// 一个没啥用的测试接口，正常一个前端传过来一个user的混杂uidStr后端解码之后
func Me(c *gin.Context) {
	// A *model.User will eventually be added to context in middleware
	//user, exists := c.Get("user")
	req := app.NewResponse(c)
	uidStr := c.Query("uid")
	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		global.Logger.Debugf(c, "Unable to extract user from request context for unknown reason: %v\n", c)
		req.ToErrorResponse(errcode.ServerError.WithDetails("This shouldn't happen, as our middleware ought to throw an error."))
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
	u.Password = "" //因为懒得写db字段，于是就这样吧
	req.ToResponse(u)
}
