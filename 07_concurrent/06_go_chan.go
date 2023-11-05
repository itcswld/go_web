package main

import "fmt"

//緩衝channel

func main() {
	//設定緩衝區大小 -- 發送方資料可放緩衝區， 所以可同時發送多個資料， 等接受方接受資料， 而非立刻收資料.  但緩衝區一滿，發送方就無法發送
	ch := make(chan int, 3) //最多3筆資料
	ch <- 1
	ch <- 2
	ch <- 3
	//ch <- 4 //channel will deadlock!
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
