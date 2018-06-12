package main

import (
	"github.com/gin-gonic/gin"
	"muxin.io/chronos/controllers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"muxin.io/chronos/interceptors"
)

func main() {
	// 定义site
	site := gin.New()

	// 中间件支持
	site.Use(gin.Logger())
	site.Use(gin.Recovery())
	// Session支持
	store := cookie.NewStore([]byte("muxin.io"))
	site.Use(sessions.Sessions("session", store))

	// 分组定义
	normal := site.Group("/api")
	{
		normal.POST("/login", controllers.DoLogin)
		normal.GET("/test2", controllers.Test)
	}
	login := site.Group("/api", interceptors.LoginInterceptor())
	{
		login.GET("/test", controllers.Test)
	}

	site.Run(":8080")
}
