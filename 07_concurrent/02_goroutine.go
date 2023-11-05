package main

//多個goroutine通訊問題
import (
	"fmt"
)

func add(a, b int) {
	c := a + b
	fmt.Println(c)
}

func main() {
	//Parallel 多個cpu下可平行
	//同時進行5個goroutine
	for i := 0; i < 5; i++ {
		go add(i, i) //被呼叫的函數返回時， goroutine也自動結束， add()來不及執行就結束了
	}
}
