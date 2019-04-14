package router

import (
	"api_server/handler/sd"
	"api_server/router/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	//Middlewares
	g.Use(gin.Recovery())     //处理某些请求时可能因为程序 bug 或者其他异常情况导致程序 panic 这时候为了不影响下一次请求的调用，通过该方法恢复 api 服务器
	g.Use(middleware.NoCache) //强制浏览器不适用缓存
	g.Use(middleware.Options) //浏览器跨域 options 请求设置
	g.Use(middleware.Secure)  //一些安全设置
	g.Use(mw...)              //载入其他的中间件

	//404
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusOK, "The incorrect API route.")
	})

	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
