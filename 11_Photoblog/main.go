package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("tpl/*"))
}
func main() {
	http.HandleFunc("/", index)
	//dir to pic path
	http.Handle("/pub/", http.StripPrefix("/pub",http.FileServer(http.Dir("./pub"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
func index(w http.ResponseWriter, req *http.Request) {
	c := getCookie(w, req)
	if req.Method == http.MethodPost {
		mf, fh, err := req.FormFile("nf")
		if err != nil {
			fmt.Println(err)
		}
		defer mf.Close()
		//create sha for file Name
		ext := strings.Split(fh.Filename, ".")[1]
		sha := sha1.New()
		io.Copy(sha, mf)
		fname := fmt.Sprintf("%x", sha.Sum(nil)) + "." + ext
		//create new file
		dir, err := os.Getwd()//wd = working directiriy
		if err != nil {
			fmt.Println(err)
		}
		path := filepath.Join(dir, "pub", "pic", fname)
		nf, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}
		defer nf.Close()
		//Copy
		mf.Seek(0,0)
		io.Copy(nf, mf)
		//add filename to user's Cookie
		c = appendValue(w, c, fname)
	}
	data := strings.Split(c.Value, "|")
	tpl.ExecuteTemplate(w, "index.html", data[1:])
}
func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	c , err := req.Cookie("session")
	if err != nil {
		if err != nil {
			c = &http.Cookie {
				Name: "session",
				Value: uuid.NewV4().String(),
			}
			http.SetCookie(w, c)
		}
	}
	return c
}
func appendValue(w http.ResponseWriter, c *http.Cookie, fname string) *http.Cookie {
	v := c.Value
	if !strings.Contains(v, fname) {
		v += "|" + fname
	}
	c.Value = v
	http.SetCookie(w, c)
	return c
}
