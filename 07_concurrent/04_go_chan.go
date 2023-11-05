package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 6; i <= 8; i++ {
			fmt.Println("start goroutine")
			ch <- i
			fmt.Println("end goroutine")
			time.Sleep(time.Second) //每次發送完時等待時間
		}
	}()

	for receive := range ch {
		fmt.Println("wait goroutine")
		fmt.Println(receive)
		if receive == 8 {
			break
		}
	}
	fmt.Println("exit")
}
