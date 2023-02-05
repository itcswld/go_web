package main

import (
	"log"
	"os"
	"text/template"
)

func main(){
	tpl := template.Must(template.ParseGlob("templates/*"))
	err := tpl.ExecuteTemplate(os.Stdout,"index.html","enjoy")
	if err != nil{
		log.Fatalln(err.Error())
	}
}
