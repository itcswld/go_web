package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//1.open
	f, err := os.Open("/Users/eve/Developer/MSSQL/LC_WMS/MLLCWMS_TABLE.sql")
	if err != nil {
		fmt.Println("Open() fail: ", f)
	}
	defer f.Close()
	//2.read --緩衝方式讀取
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		//EOF is the error returned by Read when no more input is available
		if err == io.EOF {
			break
		}
		fmt.Print(line)
	}
}
