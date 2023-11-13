/*make a crawer or spider
1.target: a website
2.爬： get all content
3.取： get ride off unwanted data
4.handle : save and use the data
*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

// 2.爬： get all content
func Get(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err != nil {
		err = err1
	}
	defer resp.Body.Close()
	//get content
	buf := make([]byte, 4*1024)
	for {
		n, err := resp.Body.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("no more input is available 讀取完畢")
				break
			} else {
				fmt.Println("body read err : ", err.Error())
			}
		}
		result += string(buf[:n])
	}
	return
}

// 3.迴圈不同的頁面， 將獲取的頁面內容分別儲存到對應的檔案
func SpiderPage(i int, page chan<- int) {
	//1.target: a website
	url := "https://github.com/search?q=golang&type=repositories&p=" + strconv.Itoa((i))
	//2.爬
	fmt.Println("爬取：", url)
	body, err := Get(url)
	if err != nil {
		fmt.Println("http err : ", err)
		return
	}
	//3.建檔
	fn := "page" + strconv.Itoa(i) + ".html"
	f, err := os.Create(fn)
	if err != nil {
		fmt.Println("create file : ", err.Error())
		return
	}
	//4.存入檔案
	f.WriteString(body)
	f.Close()
	page <- i
}

// 4.每個page都執行一個goroutine
func Run(strat, end int) {
	fmt.Printf("get page from %d to %d \n", strat, end)
	//使用channel以避免爬蟲還沒結束迴圈就結束
	page := make(chan int)
	for i := strat; i <= end; i++ {
		go SpiderPage(i, page)
	}
	for i := strat; i <= end; i++ {
		fmt.Printf("page %d finised \n", <-page)
	}
}

func main() {
	var strat, end int
	fmt.Printf("pls enter start page : ")
	fmt.Scan(&strat)
	fmt.Printf("pls enter end page : ")
	fmt.Scan(&end)
	Run(strat, end)
}
