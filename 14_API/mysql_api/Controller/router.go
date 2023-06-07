package router

import (
	"net/http"

	httpHandler "api.com/Controller/httpHandler"
)

func Mux() *http.ServeMux{
		//Mux
		mux := http.NewServeMux()
		//request
		mux.HandleFunc("/",httpHandler.Handler())//指定該路徑(router)"/get", registers此路徑func
		mux.HandleFunc("/crud/",httpHandler.CRUD())
		return mux
}
//https://go-api-2022.herokuapp.com/crud/
