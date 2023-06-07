//request
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://"
	json := "{\"id\":1}"
	rsp, err := http.Post(url, "application/x-www-urlendcoded", bytes.NewBuffer([]byte(json)))
	if err != nil {
		fmt.Println("err", err)
	}
	b, err := ioutil.ReadAll(rsp.Body)
	fmt.Println(string(b))
}
