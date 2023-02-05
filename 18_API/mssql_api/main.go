package main

import (
	"net/http"
	"os"
	"log"

	"wms.api/model"
	"wms.api/route"

	"github.com/ichtrojan/thoth"
	"github.com/joho/godotenv"
)

func main(){
	lg,_:= thoth.Init("log")
	if err := godotenv.Load(); err != nil{
		model.LogMsg(lg,".env file not found.")
	}
	port, exist := os.LookupEnv("PORT")
	if !exist{
		model.LogMsg(lg,"PORT not set.")
	}
	log.Fatal(http.ListenAndServe(":"+port, route.Init()))
}
