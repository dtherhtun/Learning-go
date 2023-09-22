package main

import (
	"context"
	"errors"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"github.com/dtherhtun/Learning-go/grpc/rides/pb"
)

type Rides struct {
	pb.UnimplementedRidesServer
}

func main() {
	addr := ":9292"

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("error: can't listen - %s\n", err)
	}

	srv := createServer()
	log.Printf("info: server ready on %s\n", addr)
	if err = srv.Serve(lis); err != nil {
		log.Fatalf("error: can't serve - %s\n", err)
	}
}

func createServer() *grpc.Server {
	srv := grpc.NewServer(grpc.UnaryInterceptor(timingInterceptor))

	var u Rides
	pb.RegisterRidesServer(srv, &u)
	reflection.Register(srv)

	return srv
}

func timingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	defer func() {
		duration := time.Since(start)
		log.Printf("info: %s took %v", info.FullMethod, duration)
	}()

	return handler(ctx, req)
}

func (r *Rides) Start(ctx context.Context, req *pb.StartRequest) (*pb.StartResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "no metadata")
	}
	log.Printf("info: api_key %s", md["api_key"])
	// TODO: validate req
	resp := pb.StartResponse{
		Id: req.Id,
	}

	// TODO: work
	return &resp, nil
}

func (r *Rides) End(ctx context.Context, req *pb.EndRequest) (*pb.EndResponse, error) {
	// TODO: validate req
	resp := pb.EndResponse{
		Id: req.Id,
	}

	// TODO: work
	return &resp, nil
}

func (r *Rides) Location(stream pb.Rides_LocationServer) error {
	count := int64(0)
	driverID := ""
	for {
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return status.Errorf(codes.Internal, "can't read")
		}
		// TODO: Update DB
		driverID = req.DriverId
		count++
	}
	resp := pb.LocationResponse{
		DriverId: driverID,
		Count:    count,
	}
	return stream.SendAndClose(&resp)
}
