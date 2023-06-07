package main

import (
	"html/template"
	"log"
	"os"
)

type Tshirt struct {
	Name   string
	Size   []string
	IsSold bool
}

func main() {
	s1 := []string{"S", "M", "L", "XL"}
	s2 := []string{"S", "M"}
	tpl := `
	{{if .IsSold}}
		{{range $i, $v := .Size}}
		{{$i}},{{$.Name}}:{{$v}}
		{{end}}
	{{else}}
		{{range .Size}}
		{{.}}
		{{end}}
	{{end}}
	`
	T := []Tshirt{
		{"Nike", s1, true},
		{"Addidas", s2, false},
	}
	t := template.Must(template.New("range").Parse(tpl))
	for _, ts := range T {
		err := t.Execute(os.Stdout, ts)
		if err != nil {
			log.Println("err:", err.Error())
		}
	}
}
