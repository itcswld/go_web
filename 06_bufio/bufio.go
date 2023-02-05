package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main(){
	s := "May there always be angels to watch over you\n To guide you each step of the way \n To guard you and keep you safe from all harm\n"
	sn := bufio.NewScanner(strings.NewReader(s))
	sn.Split(bufio.ScanRunes)
	for sn.Scan(){
		l := sn.Text()//token will be a line
		fmt.Println(l)
	}

}
