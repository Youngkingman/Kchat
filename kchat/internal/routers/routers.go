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
	} else {
		// r.Use(middleware.AccessLog())
		// r.Use(middleware.Recovery())
	}

	// r.Use(middleware.Tracing())
	// r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	r.Use(middleware.Translations())
	r.GET("/me", controller.Me)
	// article := v1.NewArticle()
	// tag := v1.NewTag()
	// upload := api.NewUpload()
	// r.GET("/debug/vars", api.Expvar)
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// r.POST("/upload/file", upload.UploadFile)
	// r.POST("/auth", api.GetAuth)
	// r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	// apiv1 := r.Group("/api/v1")
	// apiv1.Use() //middleware.JWT()
	// {
	// // 创建标签
	// apiv1.POST("/tags", tag.Create)
	// // 删除指定标签
	// apiv1.DELETE("/tags/:id", tag.Delete)
	// // 更新指定标签
	// apiv1.PUT("/tags/:id", tag.Update)
	// // 获取标签列表
	// apiv1.GET("/tags", tag.List)
	// }
	// websocket相关路由
	ws := r.Group("/ws/v1")
	ws.Use()
	{

	}
	return r
}
