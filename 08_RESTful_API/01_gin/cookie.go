package main

import (
	"gin_api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/cookie/set", controllers.SetCookiePage)
	r.GET("/cookie/read", controllers.ReadCookiePage)
	r.GET("/cookie/del", controllers.DeleteCookiePage)
	r.Run(":80")
}
