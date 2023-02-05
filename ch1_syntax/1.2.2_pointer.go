package main

import (
	"fmt"
)

func main() {
	//--pointer
	var x int
	xLoc := &x
	fmt.Println(xLoc)  //0x1400012c008
	fmt.Println(*xLoc) //0
	var z *int         //像flutter int? z;
	xLoc = z
	fmt.Println(x)  //0
	fmt.Println(z)  //nil
	fmt.Println(*z) //invalid memory address or nil pointer dereference
}
