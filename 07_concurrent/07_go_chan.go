package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := make(chan bool)
	go func() {
		time.Sleep(6)
		timeout <- true
	}()
	switch {
	case <-timeout:
		fmt.Println("timeout ch recived value.")
	default:
		fmt.Println("recived nothing.")

	}
}
