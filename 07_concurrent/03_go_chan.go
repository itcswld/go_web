package main

/*
CSP(Communicating Sequential Process) 同一個Process透過通訊來共用記憶體
channel goroutine
*/

import (
	"fmt"
	"reflect"
)

func main() {
	//make(chan Type) 創建channel同步交換資料 first in first out
	ch := make(chan string) //通訊共用記憶體
	//形成發送和接收狀態
	go func() {
		fmt.Println("Start GO")
		//發送--通道變數<-值
		ch <- "(1)阻塞模式接收"
		fmt.Println("End GO")
	}()
	fmt.Println("Wait Go")
	// (1)阻塞模式接收，變數接受後阻塞
	receive := <-ch
	fmt.Println("Exit")
	fmt.Println(receive)
	fmt.Println("receiver type : ", reflect.TypeOf(receive))
	fmt.Println(ch)

	go func() {
		//發送--通道變數<-值
		ch <- "(2)非阻塞模式接收"
	}()
	fmt.Println("Wait Go")
	//(2)非阻塞模式接收， 變數接受時不阻塞, 但會造成高cpu佔用，需配合select和計時器，以實現接受逾時檢測
	r, ok := <-ch
	if ok {
		fmt.Println(r)
	}
	//若channel沒收到資料，ok =false， 則關閉channel
	close(ch)

	go func() {
		//發送--通道變數<-值
		ch <- "(3)阻塞接收並忽略接收到的資料 "
	}()
	fmt.Println("Wait Go")
	select {
	//(3)阻塞接收並忽略接收到的資料 -- 阻塞收發
	case <-ch:
		//channel 發送成功
		fmt.Println("ch fetched value")
	}

}
