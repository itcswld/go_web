package main

import (
  "crypto/tls"
  "fmt"
  gomail "gopkg.in/gomail.v2"
)

func main() {
/*
  from := "eve.tong@lesenphants.com.tw"
  to := "eve174428@gmail.com"
  user := "eve.tong"
  pwd := "les@02"
  host := `smtp.lesenphants.com.tw`
  port := 25
*/
var (
	from = "cocodepub@gmail.com"
	to = from
	user = from
	pwd = "gm879170"
	host = `smtp.gmail.com`
	port = 587
)

  m := gomail.NewMessage()
  m.SetHeader("From", from)
  m.SetHeader("To",to)
  m.SetHeader("Subject", "Hello!")
  m.SetBody("text/html", "Hello <b>Eve</b>!")
  d := gomail.NewDialer(host, port, user, pwd)
//x509: certificate signed by unknown authority的解决方法
  d.TLSConfig = &tls.Config{InsecureSkipVerify:true}

  if err := d.DialAndSend(m); err != nil {
    panic(err)
  }
  fmt.Println("mail send successfully!")
}
