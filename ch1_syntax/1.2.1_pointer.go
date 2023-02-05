package main

import (
	"fmt"
	"reflect"
)

func main() {

	//Pointer
	/*
		which include memory address"&", and value "*",
		it depens on memory address.
		 & obtain from (a pointer) the address of
			a data item held in another location.
	*/

	m1 := 2                                              //BaseType
	ptr := &m1                                           //& memory place(地址)
	fmt.Println(ptr)                                     //0x14000122008
	fmt.Println(reflect.TypeOf(ptr))                     //*int , PointerType (var m1 *int)
	fmt.Println(*ptr)                                    //2,"*"get value
	*ptr = 8                                             //change x value
	fmt.Println(*ptr, " //type :", reflect.TypeOf(*ptr)) //int

	//--Pointer type with struct
	type Food struct {
		fluit string
	}

	var f *Food         //flutter: Food? f;
	fmt.Println(f)      //<nil>
	f = &Food{"Banana"} //init f
	fmt.Println(f.fluit)

	//--pointer demenstrate swap 2 value
	p1, p2 := "Ben", "John"
	fmt.Println(p1, p2)
	swap(&p1, &p2) //&(地址)
	fmt.Println(p1, p2)

	a, b := 1, 2
	exchange(&a, &b)
	fmt.Println(a, b)
}

//pointer
func swap(p1, p2 *string) { //交換值
	tmp := *p2
	*p2 = *p1
	*p1 = tmp
}

func exchange(a, b *int) {
	*a, *b = *b, *a
	fmt.Printf("*a type %T \n", *a)
}
