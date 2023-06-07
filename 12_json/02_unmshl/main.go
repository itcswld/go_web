package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type img struct {
	Width, Height int
	Title         string
	Thumbnail     []thumbnail
	Animated      bool
	IDs           []int
}

type thumbnail struct {
	URL    string
	Height int `json: "H"`
}

func main() {
	var data img
	j := `{
	"Width": 800,
	"Height": 600,
	"Title": "test",
	"Thumbnail" : [{"Url" : "https://google.com",
		"H" : 123
		},
		{"Url" : "https://yahoo.com",
		"H" : 123
		}],
	"Animated" : false,
	"IDs" : [1, 2, 3, 4]
	}`

	//value pointed to by v
	err := json.Unmarshal([]byte(j), &data)
	if err != nil {
		log.Fatalln("error unmarshalling", err)
	}

	fmt.Println(data)

	for i, v := range data.IDs {
		fmt.Println(i, v)
	}

	for i, v := range data.Thumbnail {
		fmt.Println(i, v.URL)
	}
}
