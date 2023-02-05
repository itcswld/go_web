package db

/*mysql
https://github.com/go-sql-driver/mysql
$go get -u github.com/go-sql-driver/mysql
*/

import (
	"database/sql"
	"fmt"
	// "log"
)

var con *sql.DB

/*heroku
mysql://bd33aaddf868c2:969b8f0c@us-cdbr-east-05.cleardb.net/heroku_9e698b958e100c1?reconnect=true
mysql://
//username
bd33aaddf868c2
:
//password
969b8f0c
@
//Hostname
us-cdbr-east-05.cleardb.net
//database
/heroku_9e698b958e100c1
?reconnect=true
*/
func Connect() *sql.DB{
	//sql.Open("mysql", "<username>:<pw>@tcp(<HOST>:<port>)/<dbname>")
	// db, err := sql.Open("mysql","root:1234@tcp(localhost:3306)")
	// db, err := sql.Open("mysql","root:1234@/TODO")//user:password@/dbname
	db, err := sql.Open("mysql","bd33aaddf868c2:969b8f0c@tcp(us-cdbr-east-05.cleardb.net:3306)/heroku_9e698b958e100c1")
	if err != nil{
		panic(err)
		// log.Fatal(err)
	}
	fmt.Println("database connected")
	con = db
	return con
}
