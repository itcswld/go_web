// 併發訊息
package main

import "fmt"

func User(user string) chan string {
	msg := make(chan string, 500)

	go func() {
		msg <- fmt.Sprintf("Hi %s, welcome back!", user)
	}()

	return msg
}

func main() {
	eve := User("Eve")
	jay := User("Jay")
	fmt.Println(<-eve)
	fmt.Println(<-jay)
}
