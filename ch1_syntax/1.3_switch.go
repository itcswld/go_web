//control
package main

import (
	"fmt"
)

func main() {
	fmt.Println("============= switch =============")
	//-- switch ---
	weekNo := []int{1, 2, 3, 4, 5, 6, 7}
	for _, v := range weekNo { //_ , same as swift "_" omit this var
		switch v {
		case 1:
			fmt.Println("Mon.")
		case 2:
			fmt.Println("Tue.")
		case 3:
			fmt.Println("Wed.")
		case 4:
			fmt.Println("Thur.")
		case 5:
			fmt.Println("Fri.")
		case 6:
			fmt.Println("Sat")
		default:
			break
		}
	}

	//(1)一分支多值
	fmt.Println("============= switch 一個分支多值=============")
	var lang = "go"
	switch lang {
	case "go", "dart":
		fmt.Println("popluar lang:", lang)
	}
	//(2)分支運算
	fmt.Println("=============  switch 分支運算 ============= ")
	var r int = 6
	switch {
	case r > 1 && r < 10:
		fmt.Println(r)
	}
}
