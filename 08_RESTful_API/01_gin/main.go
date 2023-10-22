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
	r.GET("/cookie/set", controllers.SetCookiePage)
	r.GET("/cookie/read", controllers.ReadCookiePage)
	r.GET("/cookie/del", controllers.DeleteCookiePage)
	//64MB
	r.MaxMultipartMemory = 64 << 20
	r.POST("/upload", controllers.UploadFilePage)
	r.Run(":80")
}
