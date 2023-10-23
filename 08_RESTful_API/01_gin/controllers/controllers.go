package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeV1Page(ctx *gin.Context) {

	//response Json
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello v1",
	})
}

func HomeV2Page(ctx *gin.Context) {
	//http header
	ctx.Header("Content-Type", "application/json; charset=utf-8")
	ctx.Header("site", "Eve")
	//response Json
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello v2",
	})

}

func LoginPage(ctx *gin.Context) {
	name := ctx.PostForm("name")
	pwd := ctx.PostForm("pwd")
	ctx.String(200, "username=%s, password=%s", name, pwd)
}

func DowloadPage(ctx *gin.Context) {
	ctx.File("/Users/eve/Sites/go_web/08_RESTful_API/01_gin/static/favicon-gopher.svg")
}

func UploadFilePage(ctx *gin.Context) {
	f, _ := ctx.FormFile("myfile")
	log.Println(f.Filename)
	err := ctx.SaveUploadedFile(f, fmt.Sprintf("./data/%s", f.Filename))
	if err != nil {
		ctx.String(http.StatusOK, fmt.Sprintf("%s upload fail! \n %v", f.Filename, err))
		return
	}
	ctx.String(http.StatusOK, fmt.Sprintf("%s uploaded", f.Filename))
}

func HomePageWithMW(ctx *gin.Context) {
	//regist middleware
	ex := ctx.MustGet("homepage").(string)
	log.Println(ex)
	ctx.String(http.StatusOK, ex)
}
