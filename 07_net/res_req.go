package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)
func main(){
	ln, err := net.Listen("tcp",":8080")
	if err != nil{
		log.Fatalln(err.Error())
	}
	defer ln.Close()

	for{
		conn, err := ln.Accept()
		if err != nil{
			log.Println(err.Error())
			continue
		}
		go handle(conn)
	}
}
func handle(conn net.Conn){
	defer conn.Close()
	//read
	request(conn)
	//write
	respond(conn)
}
func request(conn net.Conn){
		i := 0
		sn := bufio.NewScanner(conn)
		for sn.Scan(){
			l := sn.Text()
			fmt.Println(l)
			if i == 0{
				m := strings.Fields(l)[i]
				fmt.Println("***Method ",m,"***")
			}
			//stop when headers are done
			if l == "" {
				break
			}
			i++
		}
}
func respond(conn net.Conn){
	body := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>TCP svr</title>
	</head>
	<body>
		<strong>Hello!</strong>
	</body>
	</html>`


	//print back to connection
	fmt.Fprint(conn,"HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn,"Content-Length: %d\r\n",len(body))
	fmt.Fprint(conn,"Content-Type: text/html\r\n")
	fmt.Fprint(conn,"\r\n")
	fmt.Fprint(conn,body)
}

