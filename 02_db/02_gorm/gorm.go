package main

import (
	"fmt"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	server   = "localhost"
	port     = 1433
	user     = "sa"
	password = "E3ngin@@r"
	database = "edusys"
)

type Emp struct {
	No   int
	Name string
}

type Emp_Insert struct {
	Name string
}

func main() {

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", user, password, server, port, database)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("failed to connect database")
	}

	var result *gorm.DB

	//----- Read ------
	//select
	var emp Emp
	result = db.First(&emp, 4) // find Emp with integer primary key
	//result = db.First(&emp, "no = ?", 4) // find Emp with no 1
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Println(result.RowsAffected)
	fmt.Println(emp)

	//.select batch
	var emps []Emp
	result = db.Find(&emps, []int{4, 5}) // select * from emp where no in(4, 5)
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Println("selected rows: ", result.RowsAffected)
	fmt.Println(emps)
	//----- Update ------
	//Insert
	result = db.Exec("insert into emp(name) values('Adam'),('Eva')")
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Println("Inserted rows: ", result.RowsAffected)

	//update
	result = db.Where(Emp{No: 4}).Assign(Emp{Name: "James"}).FirstOrCreate(&emp) //update emp set name='james' where no = 4
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Println("updated rows: ", result.RowsAffected)

	//----- Delete -----
	db.Delete(&emp, []int{4, 5}) //delete from emp where no in{4, 5}
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Println("deleted rows: ", result.RowsAffected)
}
