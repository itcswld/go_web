package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//安全地被多個goroutine同時使用，
//自動創建和釋放連接，
//也會維護閒置連結的連結池
var db *sql.DB

// 定义一个初始化数据库的函数
func initDB() (err error) {
	//连接数据库
	db, err = sql.Open("mysql", "root:a123456@tcp(127.0.0.1:3306)/ch4")
	if err != nil {
		return err
	}

	//SetMaxOpenConns : default = 0
	db.SetMaxOpenConns(3)
	//SetMaxIdleConns: n<= 0, 將不保留閒置連接
	db.SetMaxIdleConns(0)

	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := initDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
}
