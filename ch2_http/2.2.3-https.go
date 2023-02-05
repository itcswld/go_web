//response-https
package main

import (
	"net/http"
)

func main() {
	handle := func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, r.Proto)
		w.Write([]byte(r.Proto))
	}
	svr := &http.Server{
		Addr:    "0.0.0.0:80",
		Handler: http.HandlerFunc(handle),
	}
	// http.HandleFunc("/", handle)
	// server.ListenAndServe()
	svr.ListenAndServeTLS("server.crt", "server.key")
}

//openssl req -newkey rsa:2048 -nodes -keyout server.key -x509 -days 365 -out server.crt
//https://0.0.0.0:80/
