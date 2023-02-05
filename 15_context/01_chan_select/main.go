package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool)

	go func()  {
		for {
			select {
			//if stop got value ~
			case <-stop:
				fmt.Println("got the stop channel.")
				return
			//else ~
			default:
				fmt.Println("still working")
				//refresh every one second
				time.Sleep(1 * time.Second)
			}
		}
	}()
	//after 5 second ~
	time.Sleep(5 * time.Second)
	fmt.Println("stop the gorutine.")
	stop <-true
	time.Sleep(1 * time.Second)
}
