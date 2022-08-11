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
		//测试接口，证明跑通了sql
		r.GET("/me", controller.Me)
		//测试接口，证明跑通了redis
		r.GET("/me-redis", controller.MeRedis)
	} else {
		// r.Use(middleware.AccessLog())
		// r.Use(middleware.Recovery())
	}

	g := r.Group("/home")
	g.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	g.Use(middleware.Translations())
	{
		g.GET("/userinfo", middleware.AuthJWT(), controller.UserInfo)
		g.POST("/signout", middleware.AuthJWT(), controller.Signout)
		g.POST("/signup", controller.Signup)
		g.POST("/signin", controller.Signin)
	}
	// websocket相关路由
	ws := r.Group("/ws/v1")
	ws.Use()
	{

	}
	return r
}
