//multiplexer (mux, servemx,router, server,http mux,...)responds to different URLs and Methods
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main(){
	ln, err := net.Listen("tcp", ":8080")
	if err != nil{
		log.Fatalln(err.Error())
	}
	defer ln.Close()

	for{
		conn, err := ln.Accept()
		if err != nil{
			log.Fatalln(err.Error())
			continue
		}
		go handle(conn)
	}

}
func handle(c net.Conn){
	defer c.Close()
	request(c)
}

func request(c net.Conn){
	i := 0
	sn := bufio.NewScanner(c)
	for sn.Scan(){
		l := sn.Text()
		fmt.Println(l)
		if i == 0{
			mux(c,l)
		}
		if l == ""{
			break
		}
		i++
	}
}
func mux(c net.Conn, l string){
	m := strings.Fields(l)[0]
	u := strings.Fields(l)[1]
	if m == "GET" && u == "/"{
		index(c)
	}else if m == "GET" && u == "/apply"{
		apply(c)
	}else if m == "POST" && u == "/applied"{
		applied(c)
	}
}
func index(c net.Conn){
	body := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>multiplexr</title>
	</head>
	<body>
		<strong>INDEX</strong><br>
		<a href="/">index</a><br>
		<a href="/apply">apply</a>
	</body>
	</html>`
	httpHandler(c,body)
}

func apply(c net.Conn){
	body := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>multiplexr</title>
	</head>
	<body>
		<strong>Apply</strong><br>
		<a href="/">index</a><br>
		<a href="/apply">apply</a><br>
		<form method="POST" action="/applied">
			<input type="submit">
		</form>
	</body>
	</html>`
	httpHandler(c,body)
}
func applied(c net.Conn){
	body := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>multiplexr</title>
	</head>
	<body>
		<strong>Applied</strong><br>
		<a href="/">index</a><br>
		<a href="/apply">apply</a>
	</body>
	</html>`
	httpHandler(c,body)
}
func httpHandler(c net.Conn, body string){
	//writes to conn
	fmt.Fprint(c,"HTTP/1.1 OK\r\n")
	fmt.Fprintf(c,"Content-Length: %d\r\n",len(body))
	fmt.Fprint(c,"Content-Type: text/html\r\n\r\n")
	fmt.Fprint(c,body)
}
