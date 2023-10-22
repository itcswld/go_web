package main

import (
	"fmt"
	"os"
)

func main() {
	//Open()
	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Open() err: ", err)
	}
	fmt.Println(f)
	//Close()
	err = f.Close()
	if err != nil {
		fmt.Println("Close() err : ", err)
	}
	//Openfile()
	f, err = os.OpenFile("./test.txt", os.O_CREATE|os.O_APPEND, 0666) //name檔名, flag打開方式, perm 許可權模式
	if err != nil {
		fmt.Println("OpenFile() err: ", err)
	}
	defer f.Close()
}
