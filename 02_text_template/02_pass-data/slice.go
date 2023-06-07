package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl := template.Must(template.ParseFiles("slice.html"))
	slice := []string{"Chi", "Eve", "Eva", "Eric"}
	err := tpl.Execute(os.Stdout, slice)
	if err != nil {
		log.Fatalln(err)
	}
}
