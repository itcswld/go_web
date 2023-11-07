// 自動增加整數
package main

import "fmt"

func GenerateInt() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i //直到channel要資料才把i加進channel
		}
	}()
	return ch
}

func main() {
	ch := GenerateInt()

	for i := 0; i < 50; i++ {
		fmt.Println(<-ch)
	}
}
