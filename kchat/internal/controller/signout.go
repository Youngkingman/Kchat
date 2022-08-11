package controller

import (
	"github.com/Youngkingman/Kchat/kchat/internal/model"
	"github.com/Youngkingman/Kchat/kchat/pkg/app"
	"github.com/Youngkingman/Kchat/kchat/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func Signout(c *gin.Context) {
	// 看jwt中间件的实现，c.Set传入的是个指针
	u := c.MustGet("user").(*model.User)
	resp := app.NewResponse(c)
	if err := model.DeleteToken(u.Email); err != nil {
		if err != nil {
			resp.ToErrorResponse(errcode.ErrorSignOutFail)
			return
		}
	}
	resp.ToResponse(gin.H{"msg": "signout successfully"})
}
