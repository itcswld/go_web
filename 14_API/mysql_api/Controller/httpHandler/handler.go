package httpHandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	jsonView "api.com/View"
)

func Handler() http.HandlerFunc{

	return func(w http.ResponseWriter, r *http.Request){
		//documentwrite
		w.Write([]byte("<h1>Hello server</h1>")) //utf-8
		fmt.Println(r.Method)//curl -X POST localhost:3000
		if r.Method == http.MethodGet{
			data := jsonView.Response{
				Code: http.StatusOK,
				Body: "hi",
			}
			json.NewEncoder(w).Encode(data)
		}
	}

}
