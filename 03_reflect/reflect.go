package main

import (
	"fmt"
	"reflect"
)

func main() {

	getObjType()
}

func getObjType(){

	tst := "string"
    fmt.Println(reflect.TypeOf(tst))
}
