package main

import (
	"context"
	"fmt"
	"log"

	pb "go_rpc/protobuf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":8078", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Dial err : %v\n", err)
	}
	defer conn.Close()

	client := pb.NewGreetingServiceClient(conn)
	//呼叫srv
	req := new(pb.Request)
	req.Name = "eve"
	resp, err := client.SendGreeting(context.Background(), req)
	if err != nil {
		log.Fatalf("resp err : %v \n", err)
	}
	fmt.Printf("Recevied %v\n", resp)
}

/*
$go run clinet.go
Recevied message:"Hello ! my friend."
*/
