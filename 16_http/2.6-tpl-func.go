//template
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
	return name + ", Go make money"
}

func sayHi(w http.ResponseWriter, r *http.Request) {
	htmlByte, err := ioutil.ReadFile("./2.6-tpl-func.html")
	if err != nil {
		fmt.Println("read html fail:", err.Error())
		return
	}
	func1 := func() string {
		return "let's learn Go together."
	}
	//way1鏈式操作：parse()之前調用func1
	tpl1, err := template.New("funcs").Funcs(template.FuncMap{"func1": func1, "welcome": Welcome}).Parse(string(htmlByte))
	if err != nil {
		fmt.Println("create tpl fail:", err.Error())
	}
	//way2FuncMap聲明， 在tpl字串符號呼叫
	funcMap := template.FuncMap{
		"Welcome": Welcome,
		"Doing":   Doing,
	}
	name := "Eve"
	tpl2, err := template.New("test").Funcs(funcMap).Parse("{{Welcome}}\n{{Doing .}}\n")
	//render
	tpl1.Execute(w, name)
	tpl2.Execute(w, name)
}

func main() {
	http.HandleFunc("/", sayHi)
	http.ListenAndServe(":80", nil)
}
