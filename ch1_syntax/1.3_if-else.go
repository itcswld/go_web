//control structures
package main

import (
	"fmt"
)

func main() {
	//if-lse
	fmt.Println("============= if-else =============")
	var isTrue *bool
	t := true
	isTrue = &t

	if isTrue == nil {
		fmt.Println("no value")
	} else if !*isTrue {
		fmt.Println("False")
	} else {
		fmt.Println("True")
	}
	//reslut:true
}
