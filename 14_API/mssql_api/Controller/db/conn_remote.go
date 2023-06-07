package sqlsvr

import (
	"database/sql"
	"fmt"
	"strings"

	"wms.api/model"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/mattn/go-adodb"
)

var (
	host = model.GetEvn("DB_HOST")
	usr = model.GetEvn("DB_USER")
	pwd = model.GetEvn("DB_PASS")
	port = model.GetEvn("DB_PORT")
)
type Mssql struct {
	*sql.DB
	dataSource string
	database   string
	windows    bool
	sa         SA
}

type SA struct {
	user   string
	passwd string
}

func (m *Mssql) Open() (err error) {
	var conf []string
	conf = append(conf, "Provider=SQLOLEDB")
	conf = append(conf, "Data Source="+m.dataSource)
	if m.windows {
		conf = append(conf, "integrated security=SSPI")
	}
	conf = append(conf, "Initial Catalog="+m.database)
	conf = append(conf, "user id="+m.sa.user)
	conf = append(conf, "password="+m.sa.passwd)

	m.DB, err = sql.Open("adodb", strings.Join(conf, ";"))
	if err != nil {
	return err
	}
	return nil
}

func ConnRemoteMssql(dbName string) Mssql{
	db := Mssql{
	dataSource: host,
	database:   dbName,
	// windwos: true 为windows身份验证，false 必须设置sa账号和密码
	windows: false,
	sa: SA{
			user:   usr,
			passwd: pwd,
		},
	}
	//connect SQL server
	err := db.Open()
	if err != nil {
		fmt.Println("sql open:", err)
	}
	return db
}
