package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func cookieHandler(method, urlpath, data string) {
	var req *http.Request
	client := &http.Client{}
	fmt.Println(urlpath)
	if data == "" {
		urlArr := strings.Split(urlpath, "?")
		if len(urlArr) == 2 {
			urlpath = urlArr[0] + "?" + url.PathEscape(urlArr[1])
		}
		req, _ = http.NewRequest(method, urlpath, nil)
	} else {
		req, _ = http.NewRequest(method, urlpath, strings.NewReader(data))
	}

	cookie := &http.Cookie{
		Name:     "new_cookie",
		Value:    "123123456789",
		HttpOnly: true,
	}
	req.AddCookie(cookie)
	req.Header.Add("new_cookie", "0987654321")

	//sends an HTTP request and returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
}

func main() {
	cookieHandler("GET", "https://api.restful-api.dev/objects?id=3&id=5&id=10", "")
}
