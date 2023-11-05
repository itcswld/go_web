package main

import "fmt"

//兩個goroutine
//不帶緩衝區的， 發送方發送資料同時必須有接收方對應接收資料

func Sum(s []int, ch chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	ch <- sum
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	ch := make(chan int)
	go Sum(s[len(s)/2:], ch)
	go Sum(s[:len(s)/2], ch)
	a, b := <-ch, <-ch
	fmt.Println("a = ", a, "; b = ", b)
}
