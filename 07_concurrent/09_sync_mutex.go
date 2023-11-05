package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex //"全域鎖“適用於不確定場景
	wait := sync.WaitGroup{}
	mutex.Lock()
	fmt.Println("LOCK")
	for i := 1; i <= 5; i++ {
		fmt.Println("i = ", i)
		wait.Add(1)
		go func(i int) {
			fmt.Println("not lock :", i)
			mutex.Lock()
			fmt.Println("lock :", i)
			time.Sleep(time.Second)
			mutex.Unlock()
			fmt.Println("unlock : ", i)
			defer wait.Done()
		}(i)
	}
	time.Sleep(time.Second)
	mutex.Unlock()
	fmt.Println("UNLOCK \n")
	wait.Wait()
}

/*result*
LOCK
i =  1
i =  2
i =  3
i =  4
i =  5
not lock : 5
not lock : 2
not lock : 4
not lock : 1
not lock : 3
UNLOCK

lock : 5
unlock :  5
lock : 2
unlock :  2
lock : 4
lock : 1
unlock :  4
unlock :  1
lock : 3
unlock :  3
eve@Eves-MacBook-Pro 07_concurrent % go run 09_sync_mutex.go
LOCK
not lock : 5
not lock : 2
not lock : 4
not lock : 1
not lock : 3
UNLOCK

lock : 5
unlock :  5
lock : 4
unlock :  4
lock : 1
unlock :  1
lock : 2
unlock :  2
lock : 3
unlock :  3
eve@Eves-MacBook-Pro 07_concurrent % go run 09_sync_mutex.go
LOCK
i =  1
i =  2
i =  3
i =  4
i =  5
not lock : 1
not lock : 3
not lock : 5
not lock : 2
not lock : 4
UNLOCK

lock : 1
unlock :  1
lock : 3
unlock :  3
lock : 5
unlock :  5
lock : 2
unlock :  2
lock : 4
unlock :  4
*/
