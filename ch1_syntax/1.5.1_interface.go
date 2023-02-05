//oop
package main

import (
	"fmt"
)

//interface 物件導向 -> 共同的部份
type Person struct {
	fname string
	lname string
}
type Agent struct {
	Person
	license bool
}
type human interface {
	speak()
}

func (p Person) speak() {
	fmt.Println("person speaking!")
}
func (s Agent) speak() {
	fmt.Println("Agent speaking")
}
func whoIsSpeaking(h human) {
	h.speak()
}
func main() {
	p := Person{
		fname: "Eve",
		lname: "Tung",
	}
	a := Agent{
		p,
		true,
	}
	whoIsSpeaking(p)
	whoIsSpeaking(a)
}
