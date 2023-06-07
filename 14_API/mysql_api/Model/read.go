package db

import (
	jsonView "api.com/View"
)

func ReadAll() ([]jsonView.PostRequest,  error){
	rows, err := con.Query("SELECT * FROM TODO")
		if err != nil{
			return nil, err
		}
		tasks := []jsonView.PostRequest{}
		for rows.Next() {
			data := jsonView.PostRequest{}
			rows.Scan(&data.Name, &data.Task)
			tasks = append(tasks, data)
		}
		return tasks, nil
}

func ReadByName(name string)([]jsonView.PostRequest, error){
	rows, err := con.Query("SELECT * FROM TODO WHERE name= ?", name)
	if err != nil{
		return nil, err
	}
	tasks := []jsonView.PostRequest{}
	for rows.Next(){
		data := jsonView.PostRequest{}
		rows.Scan(&data.Name, &data.Task)
		tasks = append(tasks, data)
	}
	return tasks, err

}
