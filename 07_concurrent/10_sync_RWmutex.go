package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var (
	rw    sync.RWMutex //一個控制goroutine的存取的讀寫鎖。 可加多個read lock or write lock, 常用於讀多於寫的次數的情境
	count int
)

func main() {
	ch := make(chan struct{}, 6)
	for i := 0; i < 3; i++ {
		//寫入開始後， 讀須等寫完才能繼續。
		go ReadCount(i, ch)
	}
	for i := 0; i < 3; i++ {
		go WriteCount(i, ch)
	}
	for i := 0; i < 6; i++ {
		<-ch
	}
}

func ReadCount(i int, ch chan struct{}) {
	fmt.Printf("goroutine %d reading start\n", i)
	rw.RLock()
	fmt.Println("reading")
	v := count
	rw.RUnlock()
	fmt.Printf("goroutine %d reading over, vaule = %d\n", i, v)
	ch <- struct{}{}
}

func WriteCount(i int, ch chan struct{}) {
	fmt.Printf("goroutine %d writing start \n", i)
	rw.Lock()
	fmt.Println("writing")
	v := rand.Intn(10)
	count = v
	rw.Unlock()
	fmt.Printf("goroutine %d writing over, new value = %d\n", i, v)
	ch <- struct{}{}
}
