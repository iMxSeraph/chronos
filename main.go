package main

import (
	"github.com/gin-gonic/gin"
	"muxin.io/chronos/controllers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"muxin.io/chronos/interceptors"
	cron2 "github.com/robfig/cron"
	"muxin.io/chronos/services/pushover"
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

	// 设定计划任务
	cron := cron2.New()
	cron.AddFunc("0 20 7 * * *", func() {
		pushover.Send(pushover.Yating, "早安哦亲爱的", "由于还没接入天气，先简单的跟你问个好嘿嘿...")
	})
	cron.Start()

	// 启动gin
	site.Run(":8080")
}
