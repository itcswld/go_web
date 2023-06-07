package main

import (
	"io"
	"net/http"
)


func main(){
	http.HandleFunc("/",dog)
	// http.Handle("/img/",http.FileServer(http.Dir("./img")))//only dir once
	http.Handle("/img/",http.StripPrefix("/img",http.FileServer(http.Dir("./img"))))
	http.ListenAndServe(":8080",nil)
}

func dog(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w,`<img src="/img/gopher.jpg">`)
}
