//實現在發送http請求時， 只有帶上指定的refer參數， 才能呼叫成功
package main

import "net/http"

type Refer struct {
	handler http.Handler
	refer   string
}

func (rf *Refer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Referer() == rf.refer {
		rf.handler.ServeHTTP(w, r)
	} else {
		w.WriteHeader(403)
	}
}
func hdl(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("handler"))
}

func hi(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))
}

func main() {
	rf := &Refer{
		handler: http.HandlerFunc(hdl),
		refer:   "www.mygolang.com",
	}
	http.HandleFunc("/", hi)
	http.ListenAndServe(":80", rf)
}

///doesnt work
