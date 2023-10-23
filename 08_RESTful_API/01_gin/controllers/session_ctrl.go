package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SessionPage(ctx *gin.Context) {
	//init
	session := sessions.Default(ctx)
	//Get()讀取Session值
	if session.Get("SessionKey") != "SessinValue" {
		session.Set("SessionKey", "SessinValue")
		session.Delete("hello")
		session.Save()
		//Clear()刪除整個session
		//session.Clear()
	}
	ctx.JSON(http.StatusOK, gin.H{
		"SessionKey": session.Get("SessionKey"),
	})

}
