//request
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//go run 2.2.1-server.go
func main() {
	res, err := http.Get("http://localhost:8080")
	if err != nil {
		fmt.Println("err", err)
	}
	body := res.Body
	byte, err := ioutil.ReadAll(body)
	fmt.Println(string(byte))
}
