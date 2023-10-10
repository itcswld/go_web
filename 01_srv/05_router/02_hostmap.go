//處理不同的二級域名
package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type HostMap map[string]http.Handler

func (h HostMap) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//分發機制： 根據domain name get 對應的router, 然後呼叫handle
	if handler := h[r.Host]; handler != nil {
		handler.ServeHTTP(w, r)
	} else {
		http.Error(w, "Forbidden", 403)
	}
}

func main() {
	rt1 := httprouter.New()
	rt1.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("route 1"))
	})
	rt2 := httprouter.New()
	rt2.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("route 2"))
	})
	//分別處理不同的二級域名
	hm := make(HostMap)
	hm["router1.localhost:8080"] = rt1
	hm["router2.localhost:8080"] = rt2

	log.Fatal(http.ListenAndServe(":8080", hm))
}
