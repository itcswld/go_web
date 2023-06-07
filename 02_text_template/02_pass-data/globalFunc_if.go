package main

import (
	"log"
	"os"
	"text/template"
)

type Wisdom struct{
		Quote string
		Name string
		Age int
		Alive bool
}

func main(){
	tpl := template.Must(template.ParseFiles("globalFunc_if.html"))

	Quotes := []Wisdom{
		{"life should be enjoyed not endured.","Gordon B Hinckley",98,false},
		{"The purpose of our lives is to be happy.","Dalai Lama",86,true},
		{"Life is what happens when you’re busy making other plans.","John Lennon",40,false},
		{"Many of life’s failures are people who did not realize how close they were to successwhen they gave up.","",0,false},
	}

	err := tpl.Execute(os.Stdout, Quotes)
	if err != nil{
		log.Fatalln(err.Error())
	}
}
