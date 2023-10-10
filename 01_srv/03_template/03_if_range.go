//if & range
package main

import (
	"fmt"
	"os"
	"text/template"
)

type Content struct {
	Contents []string
	Outside  string
	Kind     bool
}

func main() {
	range_tpl := `
	{{if .Kind}}
		{{range $i, $v := .Contents}}
		{{$i}} => {{$v}} ,{{$.Outside}}
		{{end}}
	{{else}}
		{{range .Contents}}
		{{.}}, {{$.Outside}}
		{{end}}
	{{end}}
	`
	str1 := []string{"1st range", "with index"}
	str2 := []string{"2nd range", "without index"}

	var contents = []Content{
		{str1, "1st outside content.", true},
		{str2, "2st outside content.", false},
	}

	tpl := template.Must(template.New("range").Parse(range_tpl))

	for _, c := range contents {
		err := tpl.Execute(os.Stdout, c)
		if err != nil {
			fmt.Println("tpl err : ", err)
		}
	}

}
