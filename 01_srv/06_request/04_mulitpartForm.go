package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
)

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tpl, _ := template.ParseFiles("04_mulitpartForm.html")
		tpl.Execute(w, nil)
	} else {
		r.ParseMultipartForm(1024) //從表單提取1024byte的data
		fmt.Fprintln(w, r.MultipartForm)

		fileHeader := r.MultipartForm.File["uploaded"][0] //第一個表頭檔
		//獲取檔案
		f, err := fileHeader.Open()
		if err != nil {
			fmt.Println("Open file err :", err)
		}
		//read file
		data, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Println("Read file err : ", err)
		}
		fmt.Println(w, string(data)) //&{map[] map[uploaded:[0x14000234000]]}
	}
}

func main() {
	http.HandleFunc("/file", upload)
	http.ListenAndServe(":8080", nil)
}
