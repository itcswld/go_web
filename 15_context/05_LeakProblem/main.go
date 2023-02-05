package main

import (
	"fmt"
	"time"
)

func main() {
	for n := range gen() {
		fmt.Println(n)
		if n == 5 {
			break
		}
		time.Sleep(1 * time.Second)
	}
	time.Sleep(1 * time.Minute)
}
//gen is a broken generator that will leak a goroutine.
//this goroutine will keep running, even "n" is no longer receive values.
func gen() <-chan int {
	ch := make(chan int)
	go func() {
		var n int //n = 0
		fmt.Println("n = ", n)
		for {
			ch <- n
			n++
		}
	}()
	return ch
}
