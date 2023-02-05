package main

import "fmt"

func main(){
	name := "Eve Tung"
	html := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="utf-8">
	<title>01_stdout</title>
	</head>
	<body>
	<h1> `+ name +`</h1>
	</body>
	</html>
	`
	fmt.Println(html)
	//$go run stdout.go > index.html
}
