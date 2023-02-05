package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/mshl", mshl)
	http.HandleFunc("/encd", encd)
	http.ListenAndServe(":8080", nil)
}

type person struct {
	Fname string
	Lname string
	Lang []string
}

func mshl(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "application/json")
	p1 := person{
		"Eve",
		"Tong",
		[]string{"Go", "Flutter", "swift"},
	}
	json, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(json)
}

func encd(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "application/json")
	p1 := person{
		"Eve",
		"Tong",
		[]string{"Go", "Flutter", "swift"},
	}
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println(err)
	}
}
