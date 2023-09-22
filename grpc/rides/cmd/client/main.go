package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/dtherhtun/Learning-go/grpc/rides/pb"
)

func main() {
	addr := "localhost:9292"
	creds := insecure.NewCredentials()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	conn, err := grpc.DialContext(
		ctx,
		addr,
		grpc.WithTransportCredentials(creds),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer conn.Close()

	log.Printf("info: connected to %s", addr)
	c := pb.NewRidesClient(conn)
	req := pb.StartRequest{
		Id:       "47a74960d6204a52b1bece53221eb458",
		DriverId: "007",
		Location: &pb.Location{
			Lat: 51.4871871,
			Lng: -0.1266743,
		},
		PassengerIds: []string{"M", "Q"},
		Time:         timestamppb.Now(),
		Type:         pb.RideType_POOL,
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	ctx = metadata.AppendToOutgoingContext(ctx, "api_key", "s3cr3t")
	resp, err := c.Start(ctx, &req)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(resp)

	ereq := pb.EndRequest{
		Id:       req.Id,
		Time:     timestamppb.Now(),
		Distance: 1.3,
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	eresp, err := c.End(ctx, &ereq)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Println(eresp)

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := c.Location(ctx)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	lreq := pb.LocationRequest{
		DriverId: "007",
		Location: &pb.Location{
			Lat: 51.4871871,
			Lng: -0.1266743,
		},
	}

	for i := 0.0; i < 0.03; i += 0.01 {
		lreq.Location.Lat += i
		if err := stream.Send(&lreq); err != nil {
			log.Fatalf("error: %s", err)
		}
	}

	lresp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(lresp)
}
