package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type User struct {
	Uid   int
	Name  string
	Phone string
}

//定义一个全局变量
var u User

//初始化数据库连接
func init() {
	//db, _ = sql.Open("mysql",
	//	"root:123456@tcp(127.0.0.1:3306)/chapter4")

	db, _ = sql.Open("mysql",
		"root:123456@tcp(127.0.0.1:3306)/chapter2")

	rows, err := db.Query("select `id` from `content` where id > ?", 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
	}
	println(rows)
}

/*
MysQL 前置處理
（1）普通的前置處理執行過程：
① client對Sql stmt 進行預留位置替換，得到完整的SQL stmt
② client發送完整的SQL stmt 到MysQL 伺服器端；
伺服器端執行完整的 SQL stmt，，並將結果返給用戶端。

（2）前置處理執行過程 :
① 把SQL stmt分成兩部分一一命令部分與資料部分；
② 把命令部分發送給MySQL server 做前置處理；
③ 把資料部分發送給 MysQL 伺服器端，MysQL 伺服器端對 stmt進行預留位置替換；
④ 伺服器端執行完整的 sQL stmt，並將結果返回給用戶端。

(3）為什麼要前置處理。
最佳化 MySQL伺服器重複執行 sQL stmt的間題，可以提升伺服器性能。
提前讓伺服器編譯，一次編譯多次執行，可以節省後續編譯的成本，避免 SQL 注入問題。

(4）Go 語言中的MySQL前置處理。
在Go 語言中，Prepare（）方法會將SQL 敘述發送給 MysQL 伺服器端，
返回一個準備好的狀態用於之後的查詢和命令。
返回值可以同時執行多個查詢和命令。
*/
// 预处理查询示例
func prepareQuery() {
	//① 把SQL stmt分成兩部分一一命令部分與資料部分；
	//② 把命令部分發送給MySQL server 做前置處理；
	stmt, err := db.Prepare("select uid,name,phone from `user` where uid > ?")
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	//③ 把資料部分發送給 MysQL 伺服器端，MysQL 伺服器端對 stmt進行預留位置替換；
	//④ 伺服器端執行完整的 sQL stmt，並將結果返回給用戶端。
	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	// 循环读取结果集中的数据
	for rows.Next() {
		err := rows.Scan(&u.Uid, &u.Name, &u.Phone)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("uid:%d name:%s phone:%s\n", u.Uid, u.Name, u.Phone)
	}
}

// 预处理插入示例
func prepareInsert() {
	stmt, err := db.Prepare("insert into user(username,phone) values (?,?)")
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	//一次編譯多次執行，可以節省後續編譯的成本，避免 SQL 注入問題。
	_, err = stmt.Exec("barry", 18799887766)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	_, err = stmt.Exec("jim", 18988888888)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	fmt.Println("insert success.")
}

func main() {
	prepareInsert()
}
