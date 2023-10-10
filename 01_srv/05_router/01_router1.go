package main

//package HttpRouter 補足ServerMux 無法使用變數實現url mod 匹配的缺陷，支援具名引數，以便開發Restful API
import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func HomePage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Welcome package HttpRouter"))
}

func main() {
	router := httprouter.New()
	router.GET("/", HomePage)
	log.Fatal(http.ListenAndServe(":8080", router))
}
