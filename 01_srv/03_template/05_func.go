package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

func Welcome() string {
	return "Welcome"
}

func Doing(name string) string {
	return name + ", Leaning Go Web"
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	html_f, err := ioutil.ReadFile("./05_func.html")
	if err != nil {
		fmt.Println("i/o err: ", err)
	}
	//tpl_func
	loveGo := func() string {
		return "This welcome to learn Go together."
	}

	tpl1, err := template.New("tpl1").Funcs(template.FuncMap{"Go": loveGo}).Parse(string(html_f))
	if err != nil {
		fmt.Println("tpl1 parse err : ", err)
		return
	}
	funcMap := template.FuncMap{"Welcome": Welcome, "Do": Doing}
	tpl2, err := template.New("tpl2").Funcs(funcMap).Parse("{{Welcome}}\n{{Do .}}\n")
	if err != nil {
		fmt.Println("tpl2 parse err : ", err)
		panic(err)
	}
	name := "Eve"
	tpl1.Execute(w, name)
	tpl2.Execute(w, name)
}

func main() {
	http.HandleFunc("/home", HomePage)
	http.ListenAndServe(":8080", nil)
}
