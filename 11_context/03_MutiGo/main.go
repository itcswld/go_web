package main

import (
	"context"
	"fmt"
	"time"
)

//shutdown multi-goroutine / goroutine within goroutine at once

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx, "worker01")
	go worker(ctx, "worker02")
	go worker(ctx, "worker03")

	time.Sleep(5 * time.Second)
	fmt.Println("stop the gorutine")
	cancel()
	time.Sleep(1 * time.Second)
}

func worker(ctx context.Context, name string) {
	go func(){
		for {
			select{
				//if ctx's chan got value~
				case <-ctx.Done():
					fmt.Println(name," got the stop channel.")
					return
				//else
				default:
					fmt.Println(name," is still working.")
					//refresh every once second
					time.Sleep(1 * time.Second)
			}
		}
	}()
}
