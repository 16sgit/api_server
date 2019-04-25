package router

import (
	"api_server/handler/v1/user"
	"api_server/router/middleware"

	"github.com/gin-gonic/gin"
)

func apiRouter(g *gin.Engine) {

	u := g.Group("/v1/user")
	u.Use(middleware.AuthMiddleware())
	{
		u.POST("", user.Create)       // 创建用户
		u.DELETE("/:id", user.Delete) // 删除用户
		u.PUT("/:id", user.Update)    // 更新用户
		u.GET("", user.List)          // 用户列表
		u.GET("/:username", user.Get) // 获取指定用户的详细信息
	}
}
