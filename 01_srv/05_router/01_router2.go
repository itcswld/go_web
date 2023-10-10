package main

//實現url mod 匹配的缺陷
//URL包括两种匹配模式：/user/:name精确匹配、/user/*name匹配所有的模式

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	rt := httprouter.New()
	rt.GET("/default", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("default get"))
	})
	rt.POST("/default", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("default post"))
	})

	//精確match
	rt.GET("/user/name", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Write([]byte("user name :" + p.ByName("name")))
	})

	//match all
	rt.GET("/employee/*name", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Write([]byte("employee name :" + p.ByName("name")))
	})
	http.ListenAndServe(":8080", rt)

}
