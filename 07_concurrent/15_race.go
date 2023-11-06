package main

/*race 競爭狀態檢測
如果两个以上的goroutine 在没有互相同步的情况下，试图同时读和写共享的资源，就处于相互竞争的状态
應多執行race檢測， 以避免競爭條件出現
*/

import "fmt"

// 模擬非法競爭
func main() {
	ch := make(chan bool)
	m := make(map[int]string)
	go func() {
		m[1] = "one" //conflict 1
		ch <- true
	}()
	m[2] = "two" //conflict 2
	<-ch
	for k, v := range m {
		fmt.Println(k, v)
	}
}

//go run -race 15_race.go
