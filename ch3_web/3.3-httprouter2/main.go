package main

//https://github.com/GoesToEleven/golang-web-dev/tree/master/021_third-party-serveMux
import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("tpl/*"))
}
func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "index.html", nil)
	handlerError(w, err)
}
func Hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", p.ByName("name"))
}
func apply(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "apply.html", nil)
	handlerError(w, err)
}
func httpPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "post.html", nil)
	handlerError(w, err)
}

//error
func handlerError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
func main() {
	router := httprouter.New()
	router.GET("/", index)            //type Handle func(http.ResponseWriter, *http.Request, Params)
	router.GET("/hello/:name", Hello) //localhost:8080/hello/eve
	router.GET("/apply", apply)       //localhost:8080/apply
	router.POST("/apply", httpPost)   //when localhost:8080/apply submit POST method will execute httpPost func
	http.ListenAndServe(":8080", router)
}
