//control structures
package main

import (
	"fmt"
)

func main() {

	//for loops
	fmt.Println("============= for loop =============")
	for i := 0; i < 10; i++ {
		fmt.Println(i) //1~9
	}
	//do-while
	fmt.Println("============= do-while =============")
	i := 0
	for {
		i++
		if i > 50 {
			fmt.Println("do-while:", i)
			break
		}
	}

	//for-range //swift- for...in ;flutter- forEach
	fmt.Println("============= for-range =============")
	for index, val := range []string{"apple", "banana", "orange"} {
		fmt.Println(index+1, ":", val)
	}
	//遍歷字串
	var str = "Go 加油"
	fmt.Println("============= for-range字串 =============")
	for i, v := range str {
		fmt.Println(i, ":", v)
	}

	fmt.Println("============= for-range map =============")
	//swift dictionary; flutter maps
	dic := make(map[string]interface{})
	dic["台灣"] = "Taiwan"
	dic["美國"] = "USA"
	for k, v := range dic {
		fmt.Printf("key: %s and value: %s \n", k, v)
	}
	//匿名變數"_"
	for _, v := range dic {
		fmt.Println("Eng:", v)
	}
	fmt.Println("============= for-range channel=============")
	c := make(chan int)
	//channel在遍歷時只輸出一個值
	go func() { //launch goroutine
		c <- 1 //push value to c
		c <- 2
		c <- 3
		close(c)
	}()
	//不斷接受資料直到channel被關
	for v := range c {
		fmt.Println("channel c = ", v)
	}
	fmt.Println("============= for-continue =============")
	//--- for continue
	var even []int
	//odd
	for i := 0; i < 10; i++ {
		// even
		if i%2 == 0 {
			even = append(even, i)
			continue
		} else if i == 9 {
			break
		}
		fmt.Println(i) //1,3,5,7,9
	}
	fmt.Println(even) //0~4

	fmt.Println("============= jump loop =============")
	//Tag選擇中斷那個迴圈
JumpLoop:
	for j := 0; j < 5; j++ {
		for i := 0; i < 5; i++ {
			if i > 2 {
				break JumpLoop
			}
			fmt.Println("JumpLopp:", i)
		}
	}
	var k int
JumpLoop2:
	for ; ; k++ {
		fmt.Println("k = ", k)
		if k > 10 {
			fmt.Println("JumpLoop2:", k)
			break JumpLoop2
		}
	}
	for {
		if k > 10 {
			fmt.Println("same as loop2:", k)
			break
		}
		k++
	}
	var l int
	for l <= 10 {
		l++
		fmt.Println("loop3 same as loop2:", l)
	}

}
