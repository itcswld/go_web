package httpHandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	db "api.com/Model"
	jsonView "api.com/View"
)

func CRUD() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		//Post
		if r.Method == http.MethodPost{
			//take data then save it
			data := jsonView.PostRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			//if error occured,like wrong syntax...
			if err := db.Insert(data.Name , data.Task); err != nil{ //db.Insert() will return error, it's either error message nor nil
				w.Write([]byte("error occur")) //type []byte ,anonymous var
				return
			}
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("<h1>data inserted!</h1><br>"))
			//export to Json form
			json.NewEncoder(w).Encode(data)

		//Get
		}else if r.Method == http.MethodGet{
			name := r.URL.Query().Get("name")

			var (
				data  []jsonView.PostRequest
				err  error
			)
			//no ternary in Golang
			if name != ""{
				data, err = db.ReadByName(name)
			}else{
				data, err = db.ReadAll()
			}
			if err != nil {
				w.Write([]byte(err.Error()))
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)

			todo, err :=json.MarshalIndent(data,""," ")
			fmt.Println(string(todo))
		//Delete
		}else if r.Method == http.MethodDelete{
			name := r.URL.Path[6:] //start with first string , omit the HandleFunc("pattern" string
			fmt.Println(name)
			if err := db.Delete(name); err != nil {
				w.Write([]byte(err.Error()))
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(struct{
				Status string `json:"status"`
			}{name + " is deleted!"})
		}
	}

}
