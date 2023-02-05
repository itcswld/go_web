package main

import (
	"fmt"
	"strings"
)

func main() {
	//--Types
	//--int32 vs int64
	fmt.Println("--- int32 vs int64 ---")
	var n1 int32
	var n2 int64
	fmt.Println(int64(n1) + n2) //0 , because default value is 0

	//--string
	fmt.Println("--- String ---")
	var name string
	name = "Eve"
	country := "Taiwan"
	fmt.Println(name, "is come from", country)
	fmt.Println(strings.Contains(name, country))    //false
	fmt.Println(strings.ReplaceAll(name, "e", "a")) //Eva

	//to array
	fmt.Println(strings.Split(name, ""))  //[E v e]
	fmt.Println(strings.Split(name, " ")) //[Eve]
	arrSplited := strings.Split(name, "")
	arrSplited1 := strings.Split(name, " ")
	fmt.Println(arrSplited[0])  //E
	fmt.Println(arrSplited1[0]) //Eve
}
