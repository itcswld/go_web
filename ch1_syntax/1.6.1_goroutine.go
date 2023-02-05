//1.6_concurrency
package main

import (
	"fmt"
	"time"
)

func main() {
	//run both code side by side
	go count("sleep") //"go" run this func in the  background and then continue to the next line immediately
	count("fish")
}

func count(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
