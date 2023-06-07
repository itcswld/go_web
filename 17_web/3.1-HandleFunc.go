package main

//DefaultServeMux
import (
	"io"
	"net/http"
)

func d(w http.ResponseWriter, r *http.Request){
	io.WriteString(w,"dog doggy")
}
func c(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"cat kitty")
}
func main(){
	http.HandleFunc("/dog", d)
	http.HandleFunc("/cat/",c) //localhost:8080/cat/ same as :8080/cat

	http.ListenAndServe(":8080", nil)
}
