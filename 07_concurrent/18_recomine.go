package main

import "fmt"

func compute(x int) int {
	return 1 + x
}

// 將結果發送到各自的通道
func branch(x int) chan int {
	ch := make(chan int)
	go func() {
		ch <- compute(x)
	}()
	return ch
}

func recombine(chs ...chan int) chan int {
	ch := make(chan int)

	for _, c := range chs {
		go func(c chan int) {
			ch <- <-c //複合
		}(c)
	}
	return ch
}

func main() {
	result := recombine(branch(1), branch(2), branch(3))
	for i := 0; i < 3; i++ {
		fmt.Println(<-result)
	}
}
