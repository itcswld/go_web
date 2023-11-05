// 多個讀取， 同時讀一個資料
package main

import (
	"fmt"
	"sync"
	"time"
)

var m *sync.RWMutex

func main() {
	m = new(sync.RWMutex)
	go Reading(1)
	go Reading(2)
	time.Sleep(time.Second * 2)
}

func Reading(i int) {
	fmt.Println(i, "reading start")
	m.RLock()
	fmt.Println(i, "reading")
	time.Sleep(time.Second)
	m.RUnlock()
	println(i, "reading over")
}
