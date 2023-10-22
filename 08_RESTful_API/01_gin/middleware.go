/*
middleware = 攔截http request
response週期可註冊多個middleware
每個middleware執行不同功能， 一個middleware執行完， 才能下一個middleware。
常用於：
請求限速
api介面簽名處理
許可權驗證
統一錯誤處理
*/
package main

import (
	"gin_api/controllers"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	//用Use()來使用middleware
	//logger()列印request record
	r.Use(gin.Logger())
	r.Use(SelfDefineLogger())
	//Recovery()攔截panic錯誤，to avoid sys crash
	r.Use(gin.Recovery())

	r.GET("/home", controllers.HomePageWithMW)
	r.Run(":80")

}

func SelfDefineLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()
		//設定依附在context裡的key/value
		ctx.Set("homepage", "its is middleware")
		//call next middleware
		ctx.Next()
		st := time.Since(t)
		log.Println(st)
		//查詢請求狀態
		status := ctx.Writer.Status()
		log.Println("status: ", status)
	}
}
