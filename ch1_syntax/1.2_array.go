package main

import (
	"fmt"
)

func main() {

	//array
	fmt.Println("==== Array ====")
	var arr [2]int
	arr[0] = 1
	arr[1] = 2
	arrInt := []int{1, 2, 3, 4, 5, 6}
	arrString := []string{"hello"}
	arrString = append(arrString, "World")
	fmt.Println(arr)
	fmt.Println(len(arr)) //length
	fmt.Println(arrInt, arrString)

	//Slices Type
	fmt.Println("==== Slice ====")
	var slice []int = arrInt[1:5]
	fmt.Println(slice) //2,5 #range [1]~[5-1]; alike swift 1...5 not inculde 5
	//Slices are like references to arrays
	fmt.Println("-- Slices are like \"references\" to arrays --")
	nameArr := []string{"Chi", "Eve", "Eva", "Eric"}
	fmt.Println(nameArr)
	nameSlice1 := nameArr[0:2] //Chi, Eve
	nameSlice2 := nameArr[1:3] //Eve, Eva
	fmt.Println(nameSlice1, nameSlice2)
	nameSlice1[1] = "Siang"
	fmt.Println(nameSlice1, nameSlice2, nameArr)
}
