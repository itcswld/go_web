package main

import (
	"log"
	"os"
	"text/template"
)
type Page struct{
	Title string
	Heading string
	Input string
}

func main(){
	tpl := template.Must(template.ParseFiles("tpl.html"))
	home := Page{
		Title: "Alert",
		Heading: "Nothing is escaped with text/template",
		Input: `<script>alert("Yow!");</script>`,
	}
	err := tpl.Execute(os.Stdout,home)
	if err != nil{
		log.Fatal(err)
	}
	//go run main.go > index.html
}
