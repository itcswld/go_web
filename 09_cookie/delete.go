/*
set max age or Expires,user get ten minutes without activity, without that cookie having been reset and at which point, the session is over, that is how we get rid of cookie.
Expires sets the exact time when the cookie expires, but Expires is Deprecated.
MaxAge sets how long the cookie shoud live.
*/
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/",index)
	http.HandleFunc("/set",set)
	http.HandleFunc("/read", get)
	http.HandleFunc("/delete", delete)
	http.ListenAndServe("",nil)
}
func index(w http.ResponseWriter, _ *http.Request){
	fmt.Fprintln(w, `<p><form action="/set">
	<input type="submit" value="Set Cookie">
	</form></p>`)
}

func set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "cookie#1",
		Value: "value#1",
		Path:"/",
	})
	fmt.Fprintln(w,`<p><form action="/read">
	<input value="Read" type="submit"/>
	</form></p>`)
}

func get(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("cookie#1")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	fmt.Fprintf(w,`<h1>Cookie:<br>%v</h1>
	<p><form action="/delete">
	<input value="Delete cookie" type="submit"/>
	</form></p>`,c)
}

func delete(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("cookie#1")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	c.MaxAge = -1 //delete cookie
	http.SetCookie(w, c)
	http.Redirect(w,r, "/", http.StatusSeeOther)
}
