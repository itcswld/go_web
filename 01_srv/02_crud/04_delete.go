package main

//would delete a specified pet
import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://api.restful-api.dev/objects/ff8081818a394cb8018a77a42cae452b"
	req, _ := http.NewRequest("DELETE", url, nil)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("err: ", err)
	}

	defer res.Body.Close()
	body_res, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body_res))
}
