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
10. SQL 注入與防禦
soe 注入是一種政照手殿，透過執行惡意sql stmt，任意insert，從而使攻擊者完全控制 web應用程式後端庫伺服器。
使用 sQL 注入漏洞繞過應用程式驗證 ，比如繞web page登入web身份驗證和授權頁面， 也可以繞過網頁，
直接檢索資料庫的所有丙容；還可以惡意修改、刪除和增加資料庫內容。

※ 在編寫 SQL 指令稿時，儘量不要自己拼接 SQL敘述。

針對 sQL 注入問題，常見的防奥措施有：
(1）禁止將變數直接寫入 SOL 叙述。
(2）對使用者進行分級管理，嚴格控制使用者的許可權。
(3）對使用者輸入進行檢查，確保資料登錄的安全性:
在具體檢查翰入或提交的變數時，對單引號、雙引號、冒號等字元進行轉換或過濾。
(4）對資料庫資訊進行加密。
*/
// sql注入示例
func sqlInject(name string) {
	sqlStr := fmt.Sprintf("select uid, name, phone from user where name='%s'", name)
	fmt.Printf("SQL:%s\n", sqlStr)
	ret, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("get success, affected rows:%d\n", n)
}

func main() {
	//此時用sglinjeet0 方法翰入字串可以引發SQL注入間题，範例程式：
	sqlInject("xxx' or 1=1#")
	sqlInject("xxx'union select * from user #")
	sqlInject("xxx' and (select count (*) from user) <10 #")
}
