package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetCookiePage(ctx *gin.Context) {
	ctx.SetCookie("mycookie", "myvalue", 3600, "/", "localhost", false, true)
}

func ReadCookiePage(ctx *gin.Context) {
	data, err := ctx.Cookie("mycookie")
	if err != nil {
		ctx.String(http.StatusOK, "not found!")
		return
	}
	ctx.String(http.StatusOK, data)

}

func DeleteCookiePage(ctx *gin.Context) {
	ctx.SetCookie("mycookie", "myvalue", -1, "/", "localhost", false, true)
	ctx.String(http.StatusOK, "Cookie is Deleted!")
}
