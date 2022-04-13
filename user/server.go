package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	upd "user_rpc/user/proto"
)

const (
	port = ":50051"
)

// userServer is used to implement user.SayHello.
type userServer struct{}

// SayHello implements helloworld.GreeterServer
func (s *userServer) SayHello(ctx context.Context, in *upd.HelloRequest) (*upd.HelloReply, error) {
	return &upd.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *userServer) Run() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	svc := grpc.NewServer()
	upd.RegisterGreeterServer(svc, s)

	// Register reflection service on gRPC server.
	reflection.Register(svc)
	if err := svc.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
func main() {
	uSvc := userServer{}
	uSvc.Run()
}
