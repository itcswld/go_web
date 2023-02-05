//1.6_concurrency
package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		for {
			c1 <- "Evey 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()
	go func() {
		for {
			c2 <- "Evey 2 seconds"
			time.Sleep(time.Second * 2)
		}
	}()
	//get one and the another
	for {
		fmt.Println(<-c1) //after receive from c1, it will block block each time waiting for c2.
		//everytime try to receive from c2. we're gonna have to wait 2 sec.
		fmt.Println(<-c2)
	}
	//it's really slowing down that the first gorouting.
}
