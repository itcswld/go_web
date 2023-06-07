package main

import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
)
const (
	cookieName = "seeion-id"
)
func main() {
	http.HandleFunc("/",index)
	http.ListenAndServe("",nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(cookieName)
	if err != nil {
		id := uuid.NewV4()
		c = &http.Cookie{
			Name:cookieName,
			Value: id.String(),
			Secure: true,//https
			HttpOnly: true,
		}
		http.SetCookie(w, c)
	}
	fmt.Fprintf(w,`<h1>Cookie:<br>%v</h1>`,c)
}
