package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://api.restful-api.dev/objects?id=3&id=5&id=10"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("err: ", err)
	}
	body := resp.Body
	bytes, err := ioutil.ReadAll(body)
	fmt.Println(string(bytes))
}
