package main

//直接讀到記憶體
import (
	"fmt"
	"os"
)

func main() {

	path := "/Users/eve/Developer/MSSQL/LC_WMS/MLLCWMS_TABLE.sql"
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("ReadFile() fail: ", err)
	}
	fmt.Printf("%v\n", content)
	fmt.Printf("%v\n", string(content))
}
