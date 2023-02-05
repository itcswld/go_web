//1.6_concurrency
package main

//channel : Communicating by sharing memory
import (
	"fmt"
	"time"
)

// Write(chan<- int)
// Read(<-chan int)
// ReadWrite(chan int)

//pinger 拍球，ponger接
func pinger(pinger <-chan int, ponger chan<- int) { // "<-chan int" read only, "chan <- int" write only
	for {
		//拍球
		<-pinger
		fmt.Println("ping")
		time.Sleep(time.Second)
		//接球
		ponger <- 0
	}
}

//ponger 拍球，pinger接
func ponger(ponger <-chan int, pinger chan<- int) {
	for {
		//拍球
		<-ponger
		fmt.Println("pong")
		time.Sleep(time.Second)
		//接球
		pinger <- 0
	}
}

func main() {
	//設定pinger & ponger 角色都可接球亦可拍球
	var ping chan int //Read&Write
	ping = make(chan int)
	pong := make(chan int)

	go pinger(ping, pong)
	go ponger(pong, ping)

	//pinger開球 (the main goroutine starts the pingpong by write into the ping channel)
	ping <- 0
	select {} //或是 for { time.Sleep(time.Second)}

}
