//oop
/*A variable of interface type can store a value of any type with a method set
that is any superset of the interface.
"interface{}" 很像swift type any
物件導向的根類型
*/
package main

import (
	"fmt"
	"reflect"
)

func anytype(any interface{}) {
	fmt.Println(any, "'s type is :", reflect.TypeOf(any))
}
func main() {
	anytype("string")                     //string 's type is : string
	anytype(1)                            //1 's type is : int
	anytype(struct{}{})                   //{} 's type is : struct {}
	anytype(struct{ name string }{"Eve"}) //{Eve} 's type is : struct { name string }

	//物件導向 -> 共同的部份
	//-- interface 物件導向...共同的部份
	c := Cat{"Kitty"}
	d := Dog{"Snoopy"}
	act(c)  //Kitty  eat fish.
	act(&d) //Snoopy eat meat.
}

//--interface 物件導向
type Cat struct {
	name string
}
type Dog struct {
	name string
}

//make common function as interface ,call same func  do (eat)  different things
type Action interface {
	eat()
}

func (c Cat) eat() {
	fmt.Println(c.name, " eat fish. ")
	fmt.Println(reflect.TypeOf(c)) //main.Cat, basicType
}
func (d *Dog) eat() {
	fmt.Println(d.name, "eat meat.")
	fmt.Println(reflect.TypeOf(d)) //*main.Dog, pointerType
}
func act(a Action) {
	a.eat()
}
