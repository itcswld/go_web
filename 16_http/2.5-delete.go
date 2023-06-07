//request
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	url := "http://tmp"
	json := strings.NewReader("{\"id\":1,}")
	req, _ := http.NewRequest("DELETE", url, json) //PUT
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("err", err.Error())
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

//get 請求頁面， 並返回頁面內容
//head 獲取表頭
//post 上傳
//put client向svr傳的資料取代指定文件的內容
//delete
//options 允許client看svr的效能
//connect 把svr當跳板，讓svr代替client存取其他網頁
//trace 回應svr收到的請求， 測試或診斷用
