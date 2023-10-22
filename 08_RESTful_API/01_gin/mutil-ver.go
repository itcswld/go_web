package main

import (
	"gin_api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/home", controllers.HomeV1Page)
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/home", controllers.HomeV2Page)
	}
	r.Run(":80")
}
