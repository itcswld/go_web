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
	go sender("hi", c)
	//recive msg from channel
	//wait until recive value from channel (c <- thing), once recived from channel run next line
	for msg := range c {
		fmt.Println(msg)
	}
	/*
		for {
			msg, open := <- c
			if !open{
				break
			}
			fmt.Println(msg)
		}
	*/
}
func sender(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing //send the value of thing over the channel. pointing into the channel name will send a msg
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}
