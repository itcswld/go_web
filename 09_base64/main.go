package main

//Base64 is a binary to ASCII encoding scheme. It is designed as a way to transfer binary data reliably across channels that have limited support for different content types.
import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	s := `{
		"employees":[
			{"firstName":"John", "lastName":"Doe"},
			{"firstName":"Anna", "lastName":"Smith"},
			{"firstName":"Peter", "lastName":"Jones"}
		]
		}`
	encodeStd := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	s64_1 := base64.NewEncoding(encodeStd).EncodeToString([]byte(s))
	s64_2 := base64.StdEncoding.EncodeToString([]byte(s))
	decodeBs, err := base64.StdEncoding.DecodeString(s64_2)
	if err != nil {
		log.Fatalln("decode error:", err)
	}
	fmt.Println(len(s))
	fmt.Println(len(s64_1))
	fmt.Println(len(s64_2))
	fmt.Println(s64_1)
	fmt.Println(s64_2)
	fmt.Println(string(decodeBs))
}
