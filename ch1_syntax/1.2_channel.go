package main

import (
	"fmt"
	"reflect"
)

//channel is for syncing  goroutine
// Write(chan<- int)
// Read(<-chan int)
// ReadWrite(chan int)
//https://blog.wu-boy.com/2020/01/when-to-use-go-channel-and-goroutine/

func main() {

	//-- UnBuffered Channel --
	var ch chan int
	fmt.Println(ch)     //<nil>
	ch = make(chan int) //allocates and initializes an object of the type (slice/ map / channel)
	fmt.Println(ch)     //0xc000102060
	//或是 ch := make(chan int)
	fmt.Println(reflect.TypeOf(ch)) //chan int
	ptr := &ch
	//VS Pointer type
	fmt.Println("pointer memory location:", ptr) //0xc000124018
	fmt.Println(reflect.TypeOf(ptr))             // *chan int
	//same as "ch = make(chan int)", Pointer之值 = 初始化後的channel值
	fmt.Println("pointer value:", *ptr) //0xc000102060
	go func() {
		ch <- 1
	}()
	v := <-ch
	fmt.Println(v) // 1

	go func() {
		ch <- 2
	}()
	v = <-ch
	fmt.Println(v) // 2

	v = 3
	go func() {
		ch <- v
	}()
	fmt.Println(<-ch) //3

	//-- Buffered Channel
	c := make(chan string, 1)
	//the second paramater(buffer size) doesnt matter what int nubmer it is.
	go func() {
		c <- "process 1"
		c <- "Process 2"
		c <- "Process 3"
		close(c)
	}()
	// fmt.Println(<- ch)//process 1
	// fmt.Println(<- ch)//process 2
	for v := range c {
		fmt.Println(v)
		//process 1
		//Process 2
		//Process 3
	}

}
