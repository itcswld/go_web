package main

import (
	"io"
	"os"
	"net/http"
)

func main(){
	http.HandleFunc("/",dog)
	http.HandleFunc("/gopher.jpg",pic)
	http.ListenAndServe(":8080", nil)
}
func  dog(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","text/html; charset=utf-8")
	io.WriteString(w,`<img src="./img/ gopher.jpg">`)
}

func pic(w http.ResponseWriter, r *http.Request){
	f, err := os.Open("./img/gopher.jpg")
	if err != nil{
		http.Error(w,"file not found",404)
	}
	defer f.Close()
	io.Copy(w,f)
}
