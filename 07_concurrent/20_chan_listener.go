// 使用無緩衝channel block 主線， 等待goroutine結束。 不用再使用timeout來做逾時處理
package main

import "fmt"

func main() {
	ch, done := make(chan int), make(chan int)

	go func() {
		ch <- 100 //add data
		done <- 1 //發送完成的訊號
	}()

	for isDone := false; !isDone; {
		//監視通道的資料輸出
		select {
		case v := <-ch:
			fmt.Println(v)
		case <-done:
			fmt.Println("done")
			isDone = true //done有輸出資料，關for迴圈
		}
	}

}
