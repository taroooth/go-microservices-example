package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	g "github.com/taroooth/order/app/services/user/grpc"
	"github.com/taroooth/order/app/services/user/pb"
)

func main() {
	port := 50051
	listenPort, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, g.NewUserService())

	reflection.Register(s)

	s.Serve(listenPort)
}