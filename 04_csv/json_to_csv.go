package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Application struct
type equip struct {
    Eid   string `json:"eid"`
    User  string `json:"user"`
    Name	string `json:"name"`
}

func main() {
    // read data from file
    jsonDataFromFile, err := ioutil.ReadFile("./equip.json")

    if err != nil {
        fmt.Println(err)
    }

    // Unmarshal JSON data
    var jsonData []equip
    err = json.Unmarshal([]byte(jsonDataFromFile), &jsonData)

    if err != nil {
        fmt.Println(err)
    }

    csvFile, err := os.Create("./data.csv")

    if err != nil {
        fmt.Println(err)
    }
    defer csvFile.Close()

    writer := csv.NewWriter(csvFile)

    for _, rows := range jsonData {
        var row []string
        row = append(row, rows.Eid)
        row = append(row, rows.User)
        row = append(row, rows.Name)
        writer.Write(row)
    }

    // remember to flush!
    writer.Flush()
}
