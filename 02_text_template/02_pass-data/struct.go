package main

import (
	"log"
	"os"
	"text/template"
)

type Bus struct {
	From string
	To   string
}

func main() {
	tpl := template.Must(template.ParseFiles("struct.html"))
	b := Bus{
		From: "taoyuan",
		To:   "taipei",
	}
	err := tpl.Execute(os.Stdout, b)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
