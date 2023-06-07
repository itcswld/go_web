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
實現 MysQL 交易
（1）什麼是交易。
是一個最小的、不同可再分的工作單元， 一個交易對應一個業務。
(例如銀行帳戶轉帳業務，該業務就是一個最小的工作單元
同時這個完整的業務需要執行多次 DML (ISERT 、UPDATE、DELETE
等）敘述，共同聯合完成。
舉例來說，A轉帳給B，就需要執行兩次 UPDATE 操作。
在MysQL中只有使用了 Innodb 資料庫引擎的資料庫或表才支援交易。
交易處理用來維護資料庫的完整性，保證成批的 SQL 敘述不是全部執行行，就是全部不執行。
（2）交易的 ACID 屬性。
通常交易必須滿足 4個條件(ACID）：
	原子性（Atomicity，或稱不可分割性）、
	一致性 (Consistency ）、
	隔離性 （Isolation，又稱獨立性）、
	持久性（Durability ）。
*/
func transaction() {
	tx, err := db.Begin() // 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		fmt.Printf("begin trans failed, err:%v\n", err)
		return
	}
	_, err = tx.Exec("update user set username='james' where uid=?", 1)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql1 failed, err:%v\n", err)
		return
	}
	_, err = tx.Exec("update user set username='james' where uid=?", 3)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql2 failed, err:%v\n", err)
		return
	}
	err = tx.Commit() // 提交事务
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("commit failed, err:%v\n", err)
		return
	}
	fmt.Println("exec transaction success!")
}

// sql注入示例
/*
10. SQL 注入與防禦
soe 注入是一種政照手殿，透過執行恐發 SCL，教证。谁而路任意 souR
式后入直料麻查詢中，從而使攻路者完全控制 web應用程式後合的真的
庫伺服器。
攻聚者可以使用 sQL 注入漏洞繞過應用程式驗證 ，比如繞週登入賢遷蛋
入web身份駿證和授權頁面：也可以繞過網頁，直接檢索資料庫的所有丙容；還可以惡意修改、刪除和增加資料庫內容。
提示
在編寫 SQL 指令稿時，儘量不要自己拼接 SQL敘述。
此時用sglinjeet0 方法翰入字串可以引發SQL注入間题，範例程式奶下：
sqlInject ("xxx' or 1=1#")
sqlInject ("xxx' union select * from user 41) sqlInject ("xxx' and (select count (*) from user) <10 gr
針對 sQL 注入問題，常見的防奥措施有：
(1）禁止將變數直接寫入 SOL 叙述。
(2）對使用者進行分級管理，嚴格控制使用者的許可權。
（3）對使用者輸入進行檢查，確保資料登錄的安全性。在具體檢查翰入或提交的變數時，對單引號、雙引號、冒號等字元進行轉換或過濾。
（4）對資料庫資訊進行加密。
*/
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
	//transaction()
	//sqlInject("xxx' or 1=1#")
}
