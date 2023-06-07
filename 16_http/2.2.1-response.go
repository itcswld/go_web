//server
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	h := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello Go")
		io.WriteString(w, "hello Go \n")
		w.Write([]byte("hello Go"))

	}
	http.HandleFunc("/", h)
	log.Fatal(http.ListenAndServe(":80", nil))
}
