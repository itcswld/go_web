package main

//DefaultServerMux 預設的多工器Multiplexer 為全域變數，所以不建議使用，且三方程式會在DefaultServerMux中註冊handler， 可能會發生衝突

import (
	"fmt"
	"net/http"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Home.")
}

func AboutUs(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Us...")
}

func main() {
	//註冊處理器
	srv := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: nil, //nil 表示使用預設的Multiplexer DefaultServerMux
	}
	//註冊處理器函數
	http.HandleFunc("/home", Homepage) //http://0.0.0.0:8080/home
	http.HandleFunc("/about", AboutUs) //http://0.0.0.0:8080/about
	srv.ListenAndServe()
}
