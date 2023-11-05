package main

import (
	"sync"
	"time"
)

var rw *sync.RWMutex

func main() {
	rw = new(sync.RWMutex)
	//因讀取互斥， 所以寫入開始時， 讀取須等寫完後才能讀
	go writing(1)
	go reading(2)
	go writing(3)
	time.Sleep(time.Second * 2)

}

func reading(i int) {
	println(i, "reading start")
	rw.RLock()
	println(i, "reading")
	time.Sleep(time.Second)
	rw.RUnlock()
	println(i, "reading over")
}

func writing(i int) {
	println(i, "writing start")
	rw.Lock()
	println(i, "writing")
	time.Sleep(time.Second)
	rw.Unlock()
	println(i, "writing over")
}

/*result*
1 writing start
1 writing
2 reading start
3 writing start
1 writing over
2 reading
*/
