package main

//sync.Once.Do() 用來確保channel執行過程中只被關1次
import (
	"fmt"
	"sync"
)

var (
	wg   sync.WaitGroup
	once sync.Once
)

func f1(ch1 chan<- int) {
	//每個被等的goroutine在結束時呼叫Done()
	defer wg.Done() //減掉1個計數
	fmt.Println("1st goroutine start")
	for i := 0; i < 10; i++ {
		ch1 <- i
	}
	close(ch1)
	fmt.Println("1st goroutine exit")
}

func f2(ch1 <-chan int, ch2 chan<- int) {
	defer wg.Done() //減掉1個計數
	fmt.Println("2nd goroutine start")
	for {
		v, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- 2 * v
	}
	once.Do(func() { close(ch2) }) //確保此操作只執行1次
	fmt.Println("2nd goroutine exit")
}

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	//2次goroutine, 但實際只被執行1次
	wg.Add(2) //增加2個計數, Add() 和Done()一定要配對
	go f1(ch1)
	go f2(ch1, ch2)

	fmt.Println("waiting for all goroutine")
	wg.Wait() //等到計數=0時，goroutine結束
	fmt.Println("All goroutines are finished!")
	for v := range ch2 {
		fmt.Println(v)
	}
}

/*result*
waiting for all goroutine
1st goroutine start
1st goroutine exit
2nd goroutine start
2nd goroutine exit
All goroutines are finished!
0
2
4
6
8
10
12
14
16
18
*/
