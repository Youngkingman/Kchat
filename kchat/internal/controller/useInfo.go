package controller

import (
	"github.com/Youngkingman/Kchat/kchat/internal/model"
	"github.com/Youngkingman/Kchat/kchat/pkg/app"
	"github.com/gin-gonic/gin"
)

func UserInfo(c *gin.Context) {
	u := c.MustGet("user").(*model.User)
	resp := app.NewResponse(c)
	u.Password = ""
	resp.ToResponse(u)
}
