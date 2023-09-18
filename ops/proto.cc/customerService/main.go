package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/dtherhtun/Learning-go/ops/proto.cc/go/customer"
	"google.golang.org/grpc"
)

func main() {
	//a := customer.App{}
	//a.Initialize()
	//a.Run()

	var opts []grpc.ServerOption

	listen, err := net.Listen("tcp", "localhost:9006")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterCustomerServiceServer(grpcServer, newServer())
	fmt.Println("gRPC Server started and listening on port :9006")
	grpcServer.Serve(listen)
}
