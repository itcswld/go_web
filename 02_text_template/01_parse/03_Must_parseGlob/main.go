package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)
var tpl *template.Template //package.Type
func init(){
	tpl = template.Must(template.ParseGlob("templates/*"))
}
//parse all files
func main(){
	var one,two,three string = "one.html","two.any","three.htm"
	fmt.Printf("=== %s ===",one)
	err := tpl.ExecuteTemplate(os.Stdout, one,nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Printf("=== %s ===",two)
	err = tpl.ExecuteTemplate(os.Stdout, two,nil)
	if err != nil{
		log.Fatalln(err.Error())
	}
	fmt.Printf("=== %s ===",three)
	err = tpl.ExecuteTemplate(os.Stdout, three,nil)
	if err != nil{
		log.Fatalln(err.Error())
	}
	fmt.Printf("=== %s ===",one)
	err = tpl.Execute(os.Stdout, nil)
	if err != nil{
		log.Fatalln(err.Error())
	}

}
