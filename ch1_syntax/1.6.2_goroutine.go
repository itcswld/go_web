//1.6_concurrency
package main

import (
	"fmt"
	"time"
)

//in go when the main goroutine finishes the program exits. regardness of what any other goroutines might be doing.
func main() {
	go count("sleep") //"go" run this func in the  background and then continue to the next line immediately
	go count("fish")  //continue immediately to this line.
	//for fix "go" havent had time to do anything.
	fmt.Scanln()
}

func count(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
