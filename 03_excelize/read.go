package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main(){
	f, err := excelize.OpenFile("read.xlsx")
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	cell := f.GetCellValue("Portfolio","A1")
	fmt.Println("table:",cell)
	rows := f.GetRows("Portfolio")
	for i, row := range rows {
		if i < 1 { //fetch since row 2
			continue
		}
		for _, col := range row {
			fmt.Print(col,"\t")
		}
		fmt.Println()
	}
}
