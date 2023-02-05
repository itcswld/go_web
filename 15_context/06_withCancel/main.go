package main

import (
	"context"
	"fmt"
	"time"
)

///sloving leak problem
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() //make sure all paths cancel the context to avoid context leak
	//refresh 5 second
	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			cancel()
			break
		}
		time.Sleep(1 * time.Second)
	}
	//after for loop
	fmt.Println("the goroutine channel stopped receive value.")
	time.Sleep(1 * time.Second)
}

//to avoid leak a goroutine
func gen(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		var n int //n = 0
		fmt.Println("n = ", n)
		for {
			select {
			case <-ctx.Done():
				return
			//else keep put value to the channel
			case ch <- n:
				n++
			}
		}
	}()
	return ch
}
