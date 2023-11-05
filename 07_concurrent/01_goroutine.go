package main

import (
	"fmt"
	"time"
)

func echo(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go echo("Hello")
	echo("World")
}
