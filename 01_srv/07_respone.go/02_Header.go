//重新導向
package main

import (
	"fmt"
	"net/http"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("google", "https://google.com")
	// w.WriteHeader(http.StatusMovedPermanently)
	w.WriteHeader(301)
}

func main() {
	http.HandleFunc("/redirect", redirect)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
