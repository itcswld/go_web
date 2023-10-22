package main

//filepath.Walk() --ls -al
import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func fn_scan(path string, f os.FileInfo, err error) error {
	fmt.Printf("Scaned : %s \n", path)
	return nil
}

func main() {
	uploadDir := "upload/" + time.Now().Format("2006/01/02/")
	err := os.MkdirAll(uploadDir, 0777)
	if err != nil {
		fmt.Printf("MkdirAll() err: %v \n", err)
	}

	root := "./upload"
	err = filepath.Walk(root, fn_scan)
	if err != nil {
		fmt.Printf("Walk() err: %v", err)
	}

}
