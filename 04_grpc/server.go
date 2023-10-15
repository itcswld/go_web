package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "go_rpc/protobuf"

	"google.golang.org/grpc"
)

type GreetingServiceServer struct {
	pb.UnimplementedGreetingServiceServer
}

func (p *GreetingServiceServer) SendGreeting(ctx context.Context, req *pb.Request) (resp *pb.Response, err error) {
	name := req.Name
	if name == "eve" {
		resp = &pb.Response{
			Message: "Hello ! my friend.",
		}
	}
	err = nil
	return
}

func main() {
	port := ":8078"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("listen err: %v\n", err)
	}
	fmt.Printf("listen %s \n", port)
	grpcSrv := grpc.NewServer()
	//regist GreetingService to grpc, Embed UnimplementedGreetingServiceServer)
	pb.RegisterGreetingServiceServer(grpcSrv, &GreetingServiceServer{})
	err = grpcSrv.Serve(lis)
	if err != nil {
		log.Fatalf("server err : %v", err)
	}
}

//result listen :8078
