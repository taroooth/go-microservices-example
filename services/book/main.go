package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	g "github.com/taroooth/go-microservices-example/services/book/grpc"
	"github.com/taroooth/go-microservices-example/services/book/proto"
)

func main() {
	s := grpc.NewServer()
	proto.RegisterBookServiceServer(s, g.NewBookService())

	reflection.Register(s)

	port := 50052
	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Printf("network I/O error: %v", err)
		os.Exit(1)
	}

	log.Printf("...Waiting for localhost:%v", port)

	if err := s.Serve(conn); err != nil {
		fmt.Printf("serve error: %v", err)
		os.Exit(1)
	}
}
