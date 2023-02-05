//reflection
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var n int = 1
	fmt.Println("v init value: ", n)
	fmt.Println("v type: ", reflect.TypeOf(n))
	nv1 := reflect.ValueOf(n)
	nv2 := reflect.ValueOf(&n)
	nE := nv2.Elem()
	fmt.Println("nv1 editable ?", nv1.CanSet()) //false
	fmt.Println("v2 editable ?", nE.CanSet())   //true
	nE.SetInt(2)
	fmt.Println(n)                                                            //2
	fmt.Println(nv1.Type())                                                   //int
	fmt.Println(nv1.Kind())                                                   //int
	fmt.Println(nv1.Int())                                                    //1
	fmt.Println("n & n_v1 are same type ? ", reflect.TypeOf(n) == nv1.Type()) //true
}
