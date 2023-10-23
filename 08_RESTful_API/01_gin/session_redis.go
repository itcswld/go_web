package main

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// init redis儲存引擎
	store, err := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("password"))
	if err != nil {
		fmt.Println("redis err: ", err.Error())
	}
	r.Use(sessions.Sessions("my_session", store))

	r.GET("/session", SessionRedisPage)

	r.Run(":80")
}

func SessionRedisPage(ctx *gin.Context) {
	session := sessions.Default(ctx)
	var count int
	v := session.Get("count")
	if v == nil {
		count = 0
	} else {
		count = v.(int)
		count++
	}
	session.Set("count", count)
	session.Save()
	ctx.JSON(200, gin.H{"count": count})
}
