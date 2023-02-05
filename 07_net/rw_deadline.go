package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main(){
	ln, err := net.Listen("tcp",":8080")
	if err != nil{
		fmt.Println(err)
	}
	defer ln.Close()

	for{
		conn, err := ln.Accept()
		if err != nil{
			log.Print(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn){
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil{
		log.Println()
	}
	sn := bufio.NewScanner(conn)
	for sn.Scan(){
		ln := sn.Text()
		fmt.Println(ln)
		//write
		fmt.Fprintf(conn, "I heard you say: %s\n",ln)
	}
	defer conn.Close()
	fmt.Println("***Connection is breaked***")
}
