package main

import (
	_ "image/png"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main(){
	xf, err := excelize.OpenFile("chart.xlsx")
	if err != nil{
		panic(err)
	}
	err = xf.AddPicture("Sheet1","A5","xlsx_icon.png",`{
		"x_offset": 15,
		"y_offset": 10,
		"x_scale" : 0.1,
		"y_scale" : 0.1,
		"print_obj": true,
		"lock_aspect_ratio": false,
		"locked": false
		}`)
		if err != nil{
			panic(err)
		}
	err  = xf.SaveAs("pic.xlsx")
	if err != nil{
		panic(err)
	}
}
