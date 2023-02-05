package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type Cellmap struct{
	Cell string
	Value interface{}
}
func  importData(d []Cellmap ,nf *excelize.File){
	for  _, v := range d{
		nf.SetCellValue("sheet1",v.Cell,v.Value)
	}
}
func main(){
	cat := []Cellmap{
		{"A2","Nike"},
		{"A3", "Addias"},
		{"A4", "Gap"},
		{"B1", "S"},
		{"C1", "M"},
		{"D1", "L"},
	}
	val := []Cellmap{
		{"B2", 2},
		{"C2", 1},
		{"D2", 1},
		{"B3", 2},
		{"C3", 3},
		{"D3", 1},
		{"B4", 2},
		{"C4", 3},
		{"D4", 1},
	}
	nf := excelize.NewFile()
	importData(cat, nf)
	importData(val, nf)
	err := nf.SaveAs("create.xlsx")
	if err != nil{
		fmt.Println(err.Error())
	}
}
