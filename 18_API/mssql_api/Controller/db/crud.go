package sqlsvr

import (
	"database/sql"
	"fmt"
	"strings"

	"wms.api/model"
)
var C = make(chan *sql.Rows)

//r
func Read(db Mssql,tsql string){
	rows, err := db.Query(tsql)
	model.IfErr(err,"Read Error:")
	go data(rows,C)
}
//ud
func Update(db Mssql,tsql string){
		ud := strings.TrimSuffix(tsql," ")
		q, err := db.Query(tsql)
		model.IfErr(err,ud+"Error:")
		fmt.Println(ud+" succefully!")
		defer q.Close()
}
func data(rs *sql.Rows, c chan *sql.Rows){
	c <- rs
	// close(c)
}
