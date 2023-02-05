package main

import "fmt"

//函數執行完後， defer開始由下而上執行， 先進後執行

var lang = "Go"

func main() {
	deferPriority()
	//2
	fmt.Printf("returnStr來的lang= %s \n", returnStr())
	//3 returnStr結束後執行defer
	fmt.Printf("defer來的lang= %s \n", lang)
}

func deferPriority() {
	defer func1()
	defer func2()
	defer func3()
}

func func1() {
	fmt.Println("defer 1")
}
func func2() {
	fmt.Println("defer 2")
}
func func3() {
	fmt.Println("defer 3")
}

func returnStr() string {
	defer func() {
		lang = "dart"
	}()
	//1
	fmt.Printf("lang來的lang= %s \n", lang)
	return lang
}
