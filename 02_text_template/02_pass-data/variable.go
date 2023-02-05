package main

import (
	"log"
	"os"
	"text/template"
)

func main(){
	htmlTpl := "variable.html"
	tpl := template.Must(template.ParseFiles(htmlTpl))
	err := tpl.ExecuteTemplate(os.Stdout,htmlTpl,2022)
	if err != nil{
		log.Fatalln(err.Error())
	}
}
