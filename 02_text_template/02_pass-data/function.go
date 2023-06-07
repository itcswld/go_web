package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)
type Wisdom struct{
	Quote string
	Name string
}
var fm = template.FuncMap{
	"upcasing": strings.ToUpper,
	"top3": firstThreeWord,
}
func firstThreeWord(s string) string{
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}
func (w Wisdom) Age() int{
	return 97
}
func (w Wisdom) TakeArg(x int) int{
	return x * 2
}
func main(){
	//declare functions before parse file
	tpl := template.Must(template.New(" ").Funcs(fm).ParseFiles("function.html"))
	quotes := []Wisdom{
		{"life should be enjoyed not endured.","Gordon B Hinckley"},
		{"The purpose of our lives is to be happy.","Dalai Lama"},
		{"Life is what happens when you’re busy making other plans." ,"John Lennon"},
		{"Many of life’s failures are people who did not realize how close they were to success when they gave up.","Thomas A. Edison"},
	}
	err := tpl.ExecuteTemplate(os.Stdout,"function.html",quotes)
	if err != nil{
		log.Fatalln(err.Error())
	}
}
