package main

import (
	"bufio" //load file
	"fmt"
	"os" //open file
	"sync"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)
func main(){
	await := &sync.WaitGroup{} // = new(sync.WaitGroup)
	xlsF, err := excelize.OpenFile("create.xlsx")
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	// open
	jsonF, err := os.Open("chart.json")
	if err != nil{
		fmt.Println("open err:" + err.Error())
	}
	await.Add(2)
	// load
	var json string
	scanner := bufio.NewScanner(jsonF)
	go func ()  {
		for scanner.Scan() {
			json += scanner.Text()//load line by line
		}
		fmt.Println("loading Done!")
		await.Done()
	}()
	go func(){
		time.Sleep(time.Second * 1)
		err = xlsF.AddChart("Sheet1","F1",json)
		if err != nil{
			// panic(err)
			fmt.Printf("addChartErr: %s",err.Error())
		}
		await.Done()
	}()

	await.Wait()
	err = xlsF.SaveAs("chart.xlsx")
	if err != nil{
		fmt.Printf("Save err: %s", err)
	}
}
