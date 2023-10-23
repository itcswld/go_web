package main

import (
	"gin_api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.StaticFile("/favicon-gopher.svg", "./static/favicon-gopher.svg")
	r.GET("/home", controllers.HomeV1Page)
	r.POST("/user/login", controllers.LoginPage)
	r.GET("/download", controllers.DowloadPage)
	//64MB
	r.MaxMultipartMemory = 64 << 20
	r.POST("/upload", controllers.UploadFilePage)
	r.Run(":80")
}
