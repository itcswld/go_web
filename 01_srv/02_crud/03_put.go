package main

//would update the attributes of an existing pet, identified by a specified id.
//傳送資料取代指定文件中的內容
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	url := "https://api.restful-api.dev/objects/ff8081818a394cb8018a77a42cae452b"
	json := `{
		"name": "Apple MacBook Pro 17",
		"data": {
		   "year": 2000,
		   "price": 2050,
		   "CPU model": "apple sillicon",
		   "Hard disk size": "1 TB",
		   "color": "silver"
		}
	 }`
	body := strings.NewReader(json)
	req, _ := http.NewRequest("PUT", url, body)
	req.Header.Add("Content-Type", "application/json")

	//sends request and returns an response
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body_res, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body_res))
}
