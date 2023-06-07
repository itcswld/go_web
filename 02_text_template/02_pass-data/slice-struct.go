package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"text/template"
)

type Airplane struct{
	Airline string
	From string
	To string
}
type Bus struct {
	From string
	To string
}
type Items struct{
	Airplanes []Airplane
	Buses	[]Bus
}

func main(){
	tpl := template.Must(template.ParseFiles("slice-struct.html","slice-structs.html"))
	m := Airplane{"Mandrain","TPE","MZG"}
	c := Airplane{"China", "TPE","CA"}
	t := Airplane{"Tiger", "TPE","AUZ"}

	planes := []Airplane{
		m,c,t,
	}
	buses := []Bus{
		{"TPE","TY"},
		{"TPE","KH"},
		{"TY","TPE"},
	}

	// items := Items{
	// 	planes,
	// 	buses,
	// }

	//或是

	items := struct{
		Airplanes []Airplane
		Buses	[]Bus
	}{
		planes,
		buses,
	}

	err := tpl.Execute(os.Stdout, planes)
	if err != nil{
		log.Fatalln(err.Error())
	}
	err = tpl.ExecuteTemplate(os.Stdout,"slice-structs.html", items )
	if err != nil{
		log.Fatalln(err.Error())
	}
	fmt.Println(reflect.TypeOf(planes))
}
