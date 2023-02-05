package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

func mmDDYYYY(t time.Time) string{
	// return t.Format(time.Kitchen)
	return t.Format("01-02-2006")
}

func main(){
	fm := template.FuncMap{
		"dateFormat": mmDDYYYY,
	}
	tpl := template.Must(template.New("").Funcs(fm).ParseFiles("time.html"))
	err := tpl.ExecuteTemplate(os.Stdout,"time.html",time.Now())
	if err != nil{
		log.Fatalln(err.Error())
	}
}
