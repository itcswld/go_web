//TLS = Transport Layer Security

//--to generate key and CERTIFICATE  TLS
// openssl req -newkey rsa:2048 -nodes -keyout 03server.key -x509 -days 365 -out 03server.crt
package main

import (
	"fmt"
	"log"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	body := fmt.Sprintf("protocol: %s", r.Proto) //HTTP/2.0
	w.Write([]byte(body))
}

func main() {
	// s := new(http.Server)
	// s.Addr = ":8080"
	// s.Handler = http.HandlerFunc(handle)

	srv := http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(handle),
	}

	log.Printf("https://0.0.0.0:8080 or https://localhost:8080")
	// log.Fatal(s.ListenAndServeTLS("03server.crt", "03server.key"))
	log.Fatal(srv.ListenAndServeTLS("03server.crt", "03server.key"))

}
