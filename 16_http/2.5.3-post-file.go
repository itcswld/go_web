package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main(){
	http.HandleFunc("/",fileReader)
	http.ListenAndServe("",nil)
}

func fileReader(w http.ResponseWriter, req *http.Request){
	var fileData string
	if req.Method == http.MethodPost {
		//open
		file, header, err :=  req.FormFile("q")
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println("file: ", file, "\n header:", header, "\n err: ",err)
		defer file.Close()

		data, err := ioutil.ReadAll(file)
		if err != nil{
			http.Error(w, err.Error(),http.StatusInternalServerError)
			return
		}
		fileData = string(data)
	}
	w.Header().Set("Content-Type","text/html; chartset=utf-8")
	io.WriteString(w,`
	<form method="POST" enctype="multipart/form-data">
	<input type="file" name="q">
	<input type="submit">
	</form>
	<br>
	`+fileData)
}
