package main

import (
	"database/sql"
	"time"

	"fmt"

	"github.com/astaxie/beego"
	_ "github.com/denisenkom/go-mssqldb"
)

const (
	CONN_LIVE_TIME = 24 //连接使用时间 小时
)

var (
	db *sql.DB = nil //全局数据库连接
)

func main() {
	host := beego.AppConfig.String("localhost")
	port, err := beego.AppConfig.Int("1433")
	if err != nil {
		port = 1433
	}

	user := beego.AppConfig.String("sa")
	password := beego.AppConfig.String("E3ngin@@r")
	dbName := beego.AppConfig.String("edusys")

	connString := fmt.Sprintf("server=%s;port%d;database=%s;user id=%s;password=%s", host, port, dbName, user, password)
	db, err = sql.Open("mssql", connString)
	if err != nil {
		return
	} else {
		fmt.Println("conneted")
	}

	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(50)
	db.SetConnMaxLifetime(time.Duration(CONN_LIVE_TIME) * time.Hour)
}
