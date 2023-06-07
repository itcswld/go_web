package main

import (
	"net/http"

)

func main(){
	//returned file server redirects any request ending in "/index.html"
	http.ListenAndServe(":8080",http.FileServer(http.Dir(".")))
}
