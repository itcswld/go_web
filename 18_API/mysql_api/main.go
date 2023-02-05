package main

import (
	"log"
	"net/http"
	"os" //provides a platform-independent interface to operating system functionality

	/*mysql
	https://github.com/go-sql-driver/mysql
	$go get -u github.com/go-sql-driver/mysql
	*/
	_ "github.com/go-sql-driver/mysql"

	router "api.com/Controller"
	db "api.com/Model" //<package name> <module/path>
)
func main(){

	//get port
	port := os.Getenv("PORT") //retrieves the value of the environment variable named by the key.
	log.Print("Listening on:" , port)

	db := db.Connect()
	defer db.Close() //"defter" last func() , //無論是否拋出錯誤，defer都會執行一次. 用於預設狀況(一開始的設定、與最後的清理/還原)
	//resume
	 //local
	// http.ListenAndServe("localhost:3000",controller.Mux())

	//could driver https://go-api-2022.herokuapp.com/
	log.Fatal(http.ListenAndServe(":"+port,router.Mux()))

}
