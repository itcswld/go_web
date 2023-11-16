package main

import (
	"gin_api/controllers"

	"github.com/gin-gonic/gin"
)

// 支援多個API版本
// 分組路由
func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/home", controllers.HomeV1Page) //  url = "v1/home"
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/home", controllers.HomeV2Page) //  url = "v2/home"
	}
	r.Run(":80")
}
