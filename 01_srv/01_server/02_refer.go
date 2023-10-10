package main

import "net/http"

//實現： 發送http request時， 只有帶上指定的refer參數，該參數才能呼叫成功。 Headers的 Referer = 01010111

type Refer struct {
	handler http.Handler
	refer   string
}

//中介軟體 ： 將Refer 傳遞給 ListenAndServer()
func (ref *Refer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Referer() == ref.refer {
		ref.handler.ServeHTTP(w, r)
	} else {
		w.WriteHeader(403)
	}
}

func AboutMePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About me."))
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome Home."))
}

//跳躍web server
func main() {
	referer := &Refer{
		handler: http.HandlerFunc(AboutMePage),
		refer:   "01010111",
	}
	http.HandleFunc("/home", HomePage)
	http.ListenAndServe(":8080", referer)
}

//Headers set key Referer  value= 01010111
//localhost:8080/home
//result: About me.
