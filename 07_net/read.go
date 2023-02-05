package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main(){
	ln, err := net.Listen("tcp", ":8080")
	if err != nil{
		panic(err)
	}
	defer ln.Close()

	for{
		conn, err := ln.Accept()
		if err != nil{
			log.Println(err)
		}
		go handle(conn)
	}
}
//read
func handle(conn net.Conn){
	sn := bufio.NewScanner(conn)
	for sn.Scan(){
		ln := sn.Text()
		fmt.Println(ln)
	}
	defer conn.Close()
}
//nc localhost 8080
