package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "https://api.restful-api.dev/objects"
	body := `{
		"name": "Apple MacBook Pro 16",
		"data": {
		   "year": 2019,
		   "price": 1849.99,
		   "CPU model": "Intel Core i9",
		   "Hard disk size": "1 TB"
		}
	 }`
	contentType := "application/json"
	resp, err := http.Post(url, contentType, bytes.NewBuffer([]byte(body)))
	if err != nil {
		log.Printf("err : %s", err)
	}

	s, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(s))
}
