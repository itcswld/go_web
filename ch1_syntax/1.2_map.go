//dictionary
package main

import "fmt"

func main() {
	mi := map[string]int{
		"Eve": 33,
		"Eva": 29,
	}
	fmt.Println(mi) //map[Eva:29 Eve:33]

	var m = map[string]string{}
	m["Name"] = "eve"
	var mm = make(map[string]string)
	mm["last"] = "tong"

	name, isExist := m["Name"]
	fmt.Println("Is name Exist ?", isExist, ",", name) //true,eve
	last, isExist := m["last"]
	fmt.Println("Is last name Exist ?", isExist, last) //false
	if lastName, isExist := mm["Last"]; isExist {
		fmt.Println(m["Name"], "", lastName)
	} //; then()
}
