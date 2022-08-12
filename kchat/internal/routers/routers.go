package routers

import (
	"github.com/Youngkingman/Kchat/kchat/global"
	"github.com/Youngkingman/Kchat/kchat/internal/controller"
	"github.com/Youngkingman/Kchat/kchat/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
		// r.Use(middleware.Tracing())
		// r.Use(middleware.RateLimiter(methodLimiters))
		r.GET("/me", controller.Me)
		r.GET("/me-redis", controller.MeRedis)
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}

	home := r.Group("/home")
	home.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	home.Use(middleware.Translations())
	{
		home.GET("/userinfo", middleware.AuthJWT(), controller.UserInfo)
		home.POST("/signout", middleware.AuthJWT(), controller.Signout)
		home.POST("/signup", controller.Signup)
		home.POST("/signin", controller.Signin)
	}
	// websocket相关路由
	chat := r.Group("/chat")
	chat.Use(middleware.AuthJWT())
	{
		chat.Any("/ws", controller.ChatroomWebsocket)
		chat.GET("/roompage", controller.ChatroomHomePage)
		chat.GET("/userlist", controller.ChatroomUserList)
		chat.POST("/addchatroom")
		chat.POST("/addchatroomuser")
	}
	return r
}
