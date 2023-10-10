package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	rt := httprouter.New()

	//to generate a error page 異常統一處理
	rt.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		panic("this is an error")
	})

	//keep your server from crashing
	rt.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		//and return the http error code 500
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "err: %s", i)
	}
	log.Fatal(http.ListenAndServe(":8080", rt))
}

type Router struct {
	//是否通过重定向，给路径自定加斜杠
	RedirectTrailingSlash bool
	//是否通过重定向，自动修复路径，比如双斜杠等自动修复为单斜杠
	RedirectFixedPath bool
	//是否检测当前请求的方法被允许
	HandleMethodNotAllowed bool
	//是否自定答复OPTION请求
	HandleOPTIONS bool
	//404默认处理
	NotFound http.Handler
	//不被允许的方法默认处理
	MethodNotAllowed http.Handler
	//异常统一处理
	PanicHandler func(http.ResponseWriter, *http.Request, interface{})
}
