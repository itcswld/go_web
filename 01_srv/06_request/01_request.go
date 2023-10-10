package main

import (
	"fmt"
	"log"
	"net/http"
)

func request(w http.ResponseWriter, r *http.Request) {

	body := fmt.Sprintf(`
	host : %s
	From Addr : %s
	http Method: %s
	http UIR : %s
	URL : %s
		 Rawquery : %s
		 Fragment : %s
	Protocol: %s
		major : %d
		Minor : %d
	//Form
	ParseMultipartForm : %s
	//body length
	ContentLength: %d
	Close: %t
	`, r.Host, r.RemoteAddr,
		r.Method, r.RequestURI, r.URL.Path, r.URL.RawQuery, r.URL.Fragment,
		r.Proto, r.ProtoMajor, r.ProtoMinor,
		r.ParseMultipartForm(128),
		r.ContentLength,
		r.Close,
	)
	w.Write([]byte(body))
}

func main() {
	http.HandleFunc("/", request)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

//curl http://localhost:8080/ -X POST -d "hi=web"
//ContentLength: 6
