package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	f, err := os.Open("member.csv")
	if err != nil{
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	i := 1
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, ",")
		i ++
		fmt.Printf("NO.%d, Name: %s ,email: %s \n",i,items[0],items[2])
	}
}
