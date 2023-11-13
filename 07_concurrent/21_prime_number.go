// 篩出質數
package main

import "fmt"

func genInt() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func prime(ch chan int, n int) chan int {
	out := make(chan int)
	go func() {
		for {
			i := <-ch
			//篩掉n的倍數
			if i%n != 0 {
				out <- i
			}
		}
	}()
	return out
}

func main() {
	const max = 100
	i := genInt()
	n := <-i

	for n <= max {
		fmt.Println(n)
		i = prime(i, n)
		n = <-i
	}
}
