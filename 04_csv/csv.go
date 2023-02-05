package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os" //operating system
)

func main(){
	rows := [][]string{
		{"No.", "name","age"},
		{"1","Chi","35"},
		{"2","Eve Tung","33"},
		{"3","Eva","29"},
		{"4","Eric","28"},
	}
	fileName := "siblings.csv"
	csvfile, err := os.Create(fileName)
	if err != nil{
		log.Fatalf("Fail to create file: %s",err)
	}
	csvWrite := csv.NewWriter(csvfile)
	for _, v := range rows{
		csvWrite.Write(v)
	}
	//check error occurred
	csvWrite.Flush()
	csvfile.Close()

	fmt.Println(fileName + " is created!")
}
