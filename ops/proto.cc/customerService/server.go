package main

import (
	"context"
	"database/sql"
	"errors"
	"log"

	cust "github.com/dtherhtun/Learning-go/ops/proto.cc/customerService/customer"
	pb "github.com/dtherhtun/Learning-go/ops/proto.cc/go/customer"
	_ "github.com/mattn/go-sqlite3"
)

type CustomerServer struct {
	pb.UnimplementedCustomerServiceServer
	DB *sql.DB
}

func newServer() *CustomerServer {
	db, err := sql.Open("sqlite3", "./customer.db")
	if err != nil {
		log.Fatal(err)
	}

	s := &CustomerServer{
		DB: db,
	}

	return s
}

func (cs *CustomerServer) Sigin(ctx context.Context, request *pb.SigninRequest) (*pb.SigninResponse, error) {
	log.Println("gRPC CustomerServer Signup")

	customer := request.GetCustomer()

	var c cust.Customer
	c.ID = int(customer.GetId())
	c.Username = customer.GetUsername()
	c.Passwd = customer.GetPassword()
	c.Email = customer.GetEmail()

	if cust.ExistingUser(cs.DB, &c) {
		return nil, errors.New("user already exists")
	}

	err := cust.Signup(cs.DB, &c)
	if err != nil {
		return nil, err
	}

	return &pb.SigninResponse{
		Header: request.GetHeader(),
		Customer: &pb.Customer{
			Id:       int32(c.ID),
			Username: c.Username,
			Password: c.Passwd,
			Email:    c.Email,
		},
	}, nil
}

func (cs *CustomerServer) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Println("gRPC CustomerServer Login")

	customer := request.GetCustomer()

	var c cust.Customer
	c.ID = int(customer.GetId())
	c.Username = customer.GetUsername()
	c.Passwd = customer.GetPassword()
	c.Email = customer.GetEmail()

	_, err := cust.Login(cs.DB, &c)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		Header: request.GetHeader(),
		Customer: &pb.Customer{
			Id:       int32(c.ID),
			Username: c.Username,
			Password: c.Passwd,
			Email:    c.Email,
		},
	}, nil
}
