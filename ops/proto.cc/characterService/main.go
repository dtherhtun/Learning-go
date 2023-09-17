package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/dtherhtun/Learning-go/ops/proto.cc/go/character"
)

func main() {
	//a := character.App{}
	//a.Initialize()
	//a.Run()

	var opts []grpc.ServerOption

	listen, err := net.Listen("tcp", "localhost:9008")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterCharacterServiceServer(grpcServer, newServer())
	fmt.Println("gRPC Server started and listening on port :9008")
	grpcServer.Serve(listen)
}
