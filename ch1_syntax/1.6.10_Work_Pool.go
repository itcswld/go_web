//1.6_concurrency
/*
this is where you have a queue of work to be done.
multiple concurrent workers pulling items of the queue.
*/
package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	//multi-core processor
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)
	for j := 0; j < 100; j++ {
		fmt.Println(<-results)
	}

}

//only receive from the jobs channel; only send on the results channel
func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

//Fibonacci algorithm黃金分割數列
func fib(n int) int {
	if n <= 1 {
		return n
	}
	//從第三項開始，每項都等於前兩項之和
	return fib(n-1) + fib(n-2)
}
