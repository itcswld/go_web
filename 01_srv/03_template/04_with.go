package main

import (
	"fmt"
	"os"
	"text/template"
)

type Equip struct {
	ID       string
	Name     string
	HisU     []string
	IsEnable bool
}

func main() {
	with_tpl := `
	{{/* pipline 不為空時 */}}
	{{with .}}
		{{.ID}}
		{{.Name}}
		{{range $i, $v := .HisU}}
		the {{$i}} user is {{$v}}
		{{end}}
	{{/*pipline 為空時*/}}
	{{else}}
	{{end}}
	`

	e1 := []string{"Eve", "Chi"}
	e2 := []string{"Dan", "Tracy"}
	equips := []Equip{
		{"E01", "PC", e1, true},
		{"", "Printer", e2, false},
	}

	tpl := template.Must(template.New("with").Parse(with_tpl))
	for _, e := range equips {
		err := tpl.Execute(os.Stdout, e)
		if err != nil {
			fmt.Println("tpl err :", err)
		}
	}

}
