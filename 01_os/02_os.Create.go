package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Eve"
	html := fmt.Sprint(`
		<!DOCTYPE html>
		<html>
		<head>
		<meta charset="utf-8">
		<title>02_file</title>
		</head>
		<body>
		<h1>` + name + `</h1>
		</body>
		</html>
		`)
	//create a newFile
	newFile, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer newFile.Close()
	//paste into index.html
	io.Copy(newFile, strings.NewReader(html)) // src to dst
}
