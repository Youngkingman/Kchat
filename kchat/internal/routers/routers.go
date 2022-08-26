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
		r.Use(middleware.Cors()) //debug模式下不配置nginx只能先允许所有跨域
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
		home.POST("/login", controller.Signin) //一样的前端以后再改
	}
	// websocket相关路由
	chat := r.Group("/chat")
	chat.Use(middleware.AuthJWT())
	{
		chat.GET("/ws", controller.ChatroomWebsocket)
		// chat.GET("/roompage", controller.ChatroomHomePage)
		chat.GET("/chatterlist", controller.ChatRoomChatterList)
		chat.GET("/roomuserlist", controller.ChatroomUserList)
		chat.PUT("/addchatroomuser", controller.AddUserSToChatRoom)
		chat.GET("/checkroomaccess", controller.CheckRightForWSConn)
		// The following api are under further establishment
		chat.POST("/addchatroom", controller.AddChatRoomTest)
		chat.GET("/getallinfo", controller.GetAllChatRoomTest)
		chat.GET("/getroominfo", controller.GetChatRoomByRoomIdTest)
	}
	return r
}
