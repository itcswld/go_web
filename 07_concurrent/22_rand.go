// rand 1 or 0
package main

import "fmt"

func randNum() chan int {
	ch := make(chan int)

	go func() {
		for {
			select {
			case ch <- 0:
			case ch <- 1:
			}
		}
	}()
	return ch
}

func main() {
	gen := randNum()
	for i := 0; i < 10; i++ {
		fmt.Println(<-gen)
	}
}
