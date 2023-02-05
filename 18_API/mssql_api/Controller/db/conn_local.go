package sqlsvr

import (
	"database/sql"
	"fmt"

	"wms.api/model"

	_ "github.com/denisenkom/go-mssqldb"
)

//connect
func ConnectLocalMssql(dbName string) *sql.DB{
	connStr := fmt.Sprintf("server=%s;port=%s;usr id=%s;password=%s;database=%s;encrypt=disable;",
	host,port,usr, pwd,dbName)
	fmt.Println(connStr)
	db, err := sql.Open("mssql",connStr)
	model.IfErr(err,"Connec Error:")
	// defer connDB.Close()
	return db
}
