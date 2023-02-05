package main

import (
	"fmt"
	"math"
)

func main() {
	variables()
	noneDefineType()
}

func variables() {

	//--var
	var age int
	age = 33

	//--mutiple variables
	var n1, n2 int = 1, 2
	var (
		year = 77
		bc   = 1911
	)
	var n4, n5 = 5, "hi"
	fmt.Println(n4, n5)
	//Variables with initializers :
	month := 2
	month = n1 + n2 // month += 1
	n1, n2, n3 := 5, 5, 4
	var date int = n1 + n2 + n3
	fmt.Println("i am ", age, " years old. bron in ", year+bc, "/", month, "/", date)

	//--const
	const (
		i = true
		o = false
	)
	fmt.Println(i) //true
}

func noneDefineType() {
	//無明確類型的常數類型 ex math.Pi
	var a1 float32 = math.Pi
	var b1 complex128 = math.Pi
	//vs 被確定類型的常數
	const pi64 float64 = math.Pi
	var a2 float32 = float32(pi64)
	var b2 complex128 = complex128(pi64)
	fmt.Println("float32: a1=", a1, "vs a2=", a2)
	fmt.Println("complex128: b1=", b1, "vs b2=", b2)

}
