package main

import (
	"fmt"
	"time"
)

func foo(i int) chan int {
	ch := make(chan int)
	go func() {
		ch <- i
	}()
	return ch
}

func main() {
	ch := make(chan int)
	ch1, ch2, ch3 := foo(1), foo(3), foo(5)
	//監視各channel資料輸出， 並收集資料到channel
	go func() {
		//timeout逾時處理
		timeout := time.After(time.Second)
		for isTimeout := false; !isTimeout; {
			//取出各channel資料
			select {
			case v1 := <-ch1:
				ch <- v1
			case v2 := <-ch2:
				ch <- v2
			case v3 := <-ch3:
				ch <- v3
			case <-timeout:
				isTimeout = true //逾時
			}
		}
	}()

	//block main channel, 取出channel資料
	for i := 0; i < 3; i++ {
		fmt.Println(<-ch)
	}

}
