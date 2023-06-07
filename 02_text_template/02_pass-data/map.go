package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl := template.Must(template.ParseFiles("map.html"))
	m := map[string]int{
		"Chi":  36,
		"Eve":  33,
		"Eva":  29,
		"Eric": 28,
	}
	err := tpl.Execute(os.Stdout, m)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
