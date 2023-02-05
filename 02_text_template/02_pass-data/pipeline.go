package main

import (
	"fmt"
	"math"
	"os"
	"text/template"
)

func double(x int) int{
	return x + x
}
func square(x int) float64{
	return math.Pow(float64(x),2)
}
//√
func sqrRoot(x float64) float64{
	return math.Sqrt(x)
}

func main(){
	fm := template.FuncMap{
		"double": double,
		"square": square,
		"squareRoot": sqrRoot,
	}
	tpl := template.Must(template.New("").Funcs(fm).ParseFiles("pipeline.html"))
	err := tpl.ExecuteTemplate(os.Stdout, "pipeline.html",3)
	if err != nil{
		fmt.Println(err.Error())
	}

}
