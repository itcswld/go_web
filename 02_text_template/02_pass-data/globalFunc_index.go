package main

import (
	"log"
	"os"
	"text/template"
)

func main(){
	tpl := template.Must(template.ParseFiles("globalFunc_index.html"))
	i := []int{3,5,4,1,5,8,3}
	err := tpl.Execute(os.Stdout,i)
	if err != nil{
		log.Fatalln(err)
	}
}
