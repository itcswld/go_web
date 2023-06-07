package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

//TCP = Transmission Control Protocol
func main(){
	//a phone
	listener, err := net.Listen("tcp", ":8080")
	if err != nil{
		panic(err)
	}
	defer listener.Close()
	//loop
	for{
		//pickup phone
		conn, err := listener.Accept()//if somebody calls in , accpet  then connection
		if err != nil{
			log.Println(err)
			continue
		}
		//chatting
		io.WriteString(conn, "\n Hello this is TCP Server \n")
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintf(conn, "%v","Well, I hope!")
		//hang up
		conn.Close()
	}
	//nc localhost 8080
}
