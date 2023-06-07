package main

import (
	"log"
	"os"
	"text/template"
)

//pass data
func main(){
	tpl := template.Must(template.ParseFiles("tpl.html","variable.html"))
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.html",`data`)
	if err != nil{
		log.Fatalln(err.Error())
	}
	err = tpl.ExecuteTemplate(os.Stdout, "variable.html",2022)
	if err != nil{
		log.Fatalln(err.Error())
	}
	// fmt.Println(reflect.TypeOf(tpl))

}
