package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//靜態檔案服務： 把一個目錄託管在server上以供存取

func main() {
	rt := httprouter.New()
	rt.ServeFiles("/src/*filepath", http.Dir("./files")) //p1: The path must end with "/*filepath", p2: local path
	log.Fatal(http.ListenAndServe(":8080", rt))
}
