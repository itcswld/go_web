//oop
package main

import "fmt"

//介面類型定義了一組方法, 簡單來說就是自定義一個類型type

func main() {
	//空介面
	fmt.Println("============ 空介面 ===========")
	var v1, v2 interface{}           // var v1 , v2 any
	fmt.Println(v1 == v2, v1 == nil) //true, true

	fmt.Printf("v1, v2 type = %T \n", v1)
	//不支援相等運算
	// v1, v2 = map[string]string{}, map[string]string{}
	// fmt.Println(v1 == v2) // panic: runtime error: comparing uncomparable type map[string]string
	fmt.Println("============ 介面設定值 ===========")
	var x Num = 8
	var y NumJ = x
	var z Num = 1
	fmt.Println(y)
	fmt.Println(y.Equal(x)) //true
	fmt.Println(x.Equal(z)) //false

}

type Num int

func (x Num) Equal(i Num) bool {
	return x == i
}

type NumJ interface {
	Equal(i Num) bool
}
