package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background()) //background:root節點

	go func() {
		for {
			select {
			//if ctx's chan got value~
			case <-ctx.Done():
				fmt.Println("got the stop channel.")
				return
			//else
			default:
				fmt.Println("still working.")
				//refresh every once second
				time.Sleep(1 * time.Second)
			}
		}
	}()
	//after 5 second~
	time.Sleep(5 * time.Second)
	fmt.Println("stop the gorutine.")
	//send value to the ctx's chan
	cancel()
	time.Sleep(1 * time.Second)
}
