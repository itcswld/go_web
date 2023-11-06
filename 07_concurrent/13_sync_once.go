package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once //用於多個goroutine只執行1次的情境
	prt := func() {
		fmt.Println("print only once")
	}
	done := make(chan bool)
	for i := 0; i < 6; i++ {
		go func() {
			once.Do(prt)
			done <- true
		}()
	}
	for i := 0; i < 6; i++ {
		<-done
	}
}
