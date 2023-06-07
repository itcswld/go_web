package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

//parse files
func main(){
var one,two,three string = "one.html", "two.html", "three.html"
fmt.Printf("=== %s ===",one)
//associate
 tpl, err := template.ParseFiles(one,two)
 if err != nil{
	log.Fatalln(err.Error())
 }
 //1
 err = tpl.Execute(os.Stdout,nil)
if err != nil{
	log.Fatalln(err.Error())
}
fmt.Printf("=== %s===",three)
tpl, err = tpl.ParseFiles(three)
if err != nil{
	log.Fatalln(err.Error())
}
//3
err = tpl.ExecuteTemplate(os.Stdout, three, nil)
if err != nil{
	log.Fatalln(err.Error())
}
fmt.Printf("=== %s===",two)
//2
err = tpl.ExecuteTemplate(os.Stdout, two, nil)
if err != nil{
	log.Fatalln(err.Error())
}
fmt.Printf("=== %s===",one)
//1
err = tpl.ExecuteTemplate(os.Stdout, one, nil)
if err != nil{
	log.Fatalln(err.Error())
}
fmt.Printf("=== %s===",one)
//1
err = tpl.Execute(os.Stdout,nil)
if err != nil{
	log.Fatalln(err.Error())
}


}
