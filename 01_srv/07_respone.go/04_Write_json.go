package main

import (
	"encoding/json"
	"net/http"
)

func write(w http.ResponseWriter, r *http.Request) {

	s := struct {
		Msg string `json:"Message"`
	}{
		"Hello~ Let's Go!!!",
	}
	jsonMsg, _ := json.Marshal(s)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonMsg)
}

func main() {
	http.HandleFunc("/", write)
	http.ListenAndServe(":8080", nil)
}
