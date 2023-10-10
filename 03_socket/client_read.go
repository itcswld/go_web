package main

//透過Socket 起始交談
import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

func validateErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
	}
	os.Exit(1)
}

//讀取回應資料
func fullyRead(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		//reads data from the connection.
		n, err := conn.Read(buf[:])
		result.Write(buf[:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil

}

//client read
func main() {
	fmt.Println(os.Args)
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	//主機資訊-接受 socket client 的資料
	service := os.Args[1]
	//建立連線
	conn, err := net.Dial("tcp", service)
	validateErr(err)
	//發送請求 -傳輸資料到 server 端去
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	validateErr(err)
	//讀取回應資料
	result, err := fullyRead(conn)
	validateErr(err)

	fmt.Println(string(result))
	os.Exit(0)
}
