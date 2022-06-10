package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/joho/godotenv"
	g "github.com/taroooth/go-microservices-example/services/book/grpc"
	"github.com/taroooth/go-microservices-example/services/book/proto"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Printf("Failed to load env file: %v", err)
		log.Fatalln(err)
	}

	s := grpc.NewServer()
	proto.RegisterBookServiceServer(s, g.NewBookService())

	reflection.Register(s)

	port := os.Getenv("BOOK_SERVICE_PORT")
	conn, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
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
