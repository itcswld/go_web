package main

//os -- operating system
import (
	"fmt"
	"os"
	"time"
)

func main() {
	//Create()
	fp, err := os.Create("./newfile.txt")
	if err != nil {
		fmt.Println(fp, err)
		return
	}
	fmt.Printf("%T", fp)
	defer fp.Close() //關閉檔案， 釋放資源

	//Mkdir()-----------
	err = os.Mkdir("dir1", 0777)
	if err != nil {
		fmt.Printf("Mkdir() err: %v \n", err)
	}

	//create多級目錄來保存上傳檔案
	uploadDir := "upload/" + time.Now().Format("2006/01/02/")
	err = os.MkdirAll(uploadDir, 0777)
	if err != nil {
		fmt.Printf("MkdirAll() err: %v \n", err)
	}

	//Rename()
	old_n := "dir1"
	new_n := "dir_1"
	err = os.Rename(old_n, new_n)
	if err != nil {
		fmt.Printf("Rename() err: %v \n", err)
	}
	//Remove------------
	//Remove()
	err = os.Remove(new_n)
	if err != nil {
		fmt.Printf("Remove() err: %v \n", err)
	}
	//RemoveAll()多級目錄
	err = os.RemoveAll("upload")
	if err != nil {
		fmt.Printf("RemoveAll() err: %v \n", err)
	}
}
