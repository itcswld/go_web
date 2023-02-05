package main

import (
	"fmt"
	"os"
)

func random(ch chan int) {
	select {
	case <-ch:
		fmt.Println("random 1")
	case <-ch:
		fmt.Println("random 2")
	}
}

func main() {

	//--select
	//同一個channel select會隨機選取
	ch := make(chan int, 1)
	ch <- 1
	/*if Channel doesnt pass by value will cause a deadlock,
	for avoid crash : use "default"*/
	random(ch)
	select {
	case <-ch:
		fmt.Println("random 1")
	case <-ch:
		fmt.Println("random 2")
	//use "default" avoid crash
	default:
		fmt.Println("you've not yet sent value into Channel.")
	}

	//2 channels
	c1 := make(chan int)
	c2 := make(chan int)
	go Select(c1, c2)

	c1 <- 1
	c2 <- 0

	select {} //for{time.Sleep(time.Second)}
}

func Select(c1 chan int, c2 chan int) {
	for {
		select {
		case <-c1:
			//when c1 received value
			fmt.Println("ON")
		case <-c2:
			//when c2 received value
			fmt.Println("OFF")
			os.Exit(0)
		}
	}
}
