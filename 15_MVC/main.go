package main

import (
	"mvc/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	uc := controllers.NewUserController()

	r := httprouter.New()
	r.GET("/", uc.Index)
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}
