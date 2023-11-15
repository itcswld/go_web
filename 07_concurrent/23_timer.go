package main

import (
	"fmt"
	"time"
)

func Timer(duration time.Duration) chan bool {
	ch := make(chan bool)

	go func() {
		time.Sleep(duration)
		//time is up
		ch <- true
	}()
	return ch
}

func main() {
	timeout := Timer(time.Second * 5)

	for {
		select {
		case <-timeout:
			fmt.Println("time is up")
			return
		}
	}
}
