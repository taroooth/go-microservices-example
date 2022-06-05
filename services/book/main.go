package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	g "github.com/taroooth/order/services/book/grpc"
	"github.com/taroooth/order/services/book/proto"
)

func main() {
	port := 50052
	listenPort, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterBookServiceServer(s, g.NewBookService())

	reflection.Register(s)

	s.Serve(listenPort)
}