//handle
package main

import (
	"fmt"
	"log"
	"net/http"
)

type HandlerMsg struct {
	lang string
}

func (h HandlerMsg) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, h.lang)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/tw", HandlerMsg{"台灣"})
	mux.Handle("/en", HandlerMsg{"USA"})

	svr := &http.Server{
		Addr:    ":80",
		Handler: mux,
	}
	if err := svr.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
