package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
	cupb "user_rpc/user_client/proto"
)

const (
	address     = "localhost:50051"
	defaultName = "pangxieke"
)

var client cupb.GreeterClient

func init() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client = cupb.NewGreeterClient(conn)
}

func GetUser() (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.SayHello(ctx, &cupb.HelloRequest{Name: "test"})
	if err != nil {
		return nil, err
	}
	return r.Message, nil
}
