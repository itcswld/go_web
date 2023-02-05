package model


import (
	"errors"
	"log"
	"os"
	"fmt"

	"github.com/ichtrojan/thoth"
	_ "github.com/joho/godotenv/autoload"
)

func LogMsg(lg thoth.Config, msg string){
	lg.Log(errors.New(msg))
	log.Fatal(msg)
}
func IfErr(err error, msg string){
	if err != nil{
		log.Fatal(msg + err.Error())
	}
}

func GetEvn(key string) string{
	lgr, _:= thoth.Init("log")
	value, exist := os.LookupEnv(key)
	if !exist{
		msg := fmt.Sprintf("%s not found.",key)
		lgr.Log(errors.New(msg))
		log.Fatal(msg)
	}
	return value
}
