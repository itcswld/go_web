package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Param) {
	w.Write([]byte("hi"))
}

func main() {
	rt := httprouter.New()
	rt.GET("/", Index)
	log.Fatal(http.ListenAndServe(":80", rt))
}
