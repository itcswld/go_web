package main

import (
	"os"
	"text/template"
)

func main() {

	tpl_boolean := `
	{{.Cn}} == {{.En}} ? {{eq .Cn .En}}
	{{.Cn}} == {{.En}} || {{.Cn}} == {{.Math}}? {{eq  .Cn  .En .Math}}

	{{.Cn}} != {{.En}} ? {{ne .Cn .En}}
	{{.Cn}} < {{.Math}} ? {{lt .Cn .Math}}
	{{.Cn}} <= {{.Math}} ? {{le .Cn .Math}}
	{{.En}} > {{.Math}} ? {{gt .En .Math}}
	{{.En}} >= {{.Math}} ? {{ge .En .Math}}
	`
	Score := struct{ Cn, Math, En int }{
		50, 100, 50,
	}
	tpl := template.Must(template.New("boolean").Parse(tpl_boolean))
	tpl.Execute(os.Stdout, Score)

}
