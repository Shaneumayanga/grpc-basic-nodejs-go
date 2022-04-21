package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/Shaneumayanga/grpc-react/hello"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (s *server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	fmt.Printf("in.Password: %v\n", in.Password)
	fmt.Printf("in.Username: %v\n", in.Username)
	return &pb.LoginResponse{
		Token: "hello",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})
	log.Println("Server running on port :8000")
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
