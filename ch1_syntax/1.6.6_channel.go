//1.6_concurrency
/*
chan - communicate back to the main goroutine
its like a pipe through which you can send a msg or receive a msg.
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	//pass the channel
	c := make(chan string)
	go send("hi", c)
	//recive msg from channel
	msg := <-c //wait until recive value from channel (c <- thing), once recived from channel run next line
	fmt.Println(msg)
}
func send(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing //send the value of thing over the channel. pointing into the channel name will send a msg
		time.Sleep(time.Millisecond * 500)
	}
}
