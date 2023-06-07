package db

import "fmt"

func Insert(name, task string) error{ //method name must be capitalization
	insertQ, err := con.Query("INSERT INTO TODO VALUES(?, ?)", name, task)
	defer insertQ.Close()
	if err != nil{
		fmt.Println(err)
		return err
	}
	fmt.Println("inserted!")
	return nil
}

func Delete(name string) error{
	deleteQ, err := con.Query("DELETE FROM TODO WHERE name=?", name)
	defer deleteQ.Close()
	if err != nil{
		fmt.Println(err)
		return err
	}
	return nil
}
