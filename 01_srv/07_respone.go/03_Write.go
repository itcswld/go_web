package main

import "net/http"

func write(w http.ResponseWriter, r *http.Request) {

	html := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>get form value</title>
	</head>
	<body>
		<h1>Hello~ Let's Go!</h1>
	</body>
	</html>
	`
	w.Write([]byte(html))
}

func main() {
	http.HandleFunc("/", write)
	http.ListenAndServe(":8080", nil)
}
