//control
package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	//goto 無條件跳躍到標籤
	fmt.Println("=============  goto ============= ")
	// var isBreak bool
	for x := 0; x < 20; x++ {
		for y := 0; y < 20; y++ {
			if y == 2 {
				// isBreak = true
				// break
				goto breakTag //跳到breaTag執行
			}
		}
		// if isBreak {
		// 	break
		// }
		return
	}
	// fmt.Println("over")
breakTag:
	fmt.Println("over")

	//goto用於重複的程式
	fmt.Println("=============  goto 用於重複的程式============= ")

	err := errors.New("404")
	if err != nil {
		goto logTag
	}

	return
logTag:
	fmt.Println(err)
	fmt.Println(reflect.TypeOf(err)) //*errors.errorString
}
