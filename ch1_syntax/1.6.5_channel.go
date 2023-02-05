package main

import (
	"fmt"
)

func main() {
	c := make(chan string, 2) //wont block until the channel is full
	//sending "hello" across channel
	c <- "hello"
	c <- "World"
	//因size 2, 收到3rd就fatal error: all goroutines are asleep - deadlock!
	c <- "test"
	/*sender will block until something is ready to receive.*/
	//recive value from channel "c <-"hello"
	msg := <-c
	fmt.Println(msg)
	//recive value from channel "c <-"World"
	msg = <-c
	fmt.Println(msg)
	msg = <-c
	fmt.Println(msg)
}
