package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go work(ctx, "node 1")
	go work(ctx, "node 2")
	go work(ctx, "node 3")
	time.Sleep(time.Second * 5)
	fmt.Println("stop the gorutine")
	cancel() //ctx.Done()
	time.Sleep(time.Second * 5)

}

func work(ctx context.Context, name string) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println(name, "is stopped")
				return
			default:
				fmt.Println(name, "is working")
				time.Sleep(time.Second * 1)
			}
		}
	}()
}
