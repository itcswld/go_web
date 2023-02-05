package main

import (
	"log"
	"os"
	"text/template"
)

//copy
func main(){
	tpl, err := template.ParseFiles("tpl.html")
	if err != nil{
		log.Fatalln(err.Error())
	}
	nf, err := os.Create("index.html")
	if err != nil{
		log.Println(err.Error())
	}
	defer nf.Close()
	//tpl.html save into index.html
	err = tpl.Execute(nf,nil)
	if err != nil{
		log.Fatalln(err.Error())
	}
}
