package main

import (
	"gin_api/controllers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//創建cookie儲存引擎
	store := cookie.NewStore([]byte("pwd123"))
	//設定session middleware, ref mySession = session name  also cookie's name as well
	r.Use(sessions.Sessions("mySession", store))
	r.GET("/session", controllers.SessionPage)
	r.Run(":80")
}
