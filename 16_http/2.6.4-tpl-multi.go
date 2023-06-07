//template multipart
package main

import (
	"html/template"
	"net/http"
	"net/url"
)

func handler() http.HandlerFunc {
	tpl := template.Must(template.ParseFiles("2.6.4-tpl-multi.html"))
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			panic(err)
		}

		data := struct {
			Method string
			Host   string
			URL    *url.URL
			/*requests are large so we don't want to copy it all - use a pointer instead.
			if the data is large, use pointers to pass it around as this increases performance*/
			Submissions   url.Values //map[string][]string
			Header        http.Header
			ContentLength int64
		}{
			r.Method,
			r.Host,
			r.URL,
			r.Form,
			r.Header,
			r.ContentLength,
		}
		err = tpl.Execute(w, data)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	http.ListenAndServe(":8080", handler())
}
