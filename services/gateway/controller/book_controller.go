package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/taroooth/go-microservices-example/services/book/proto"
	"google.golang.org/grpc"
)

func GetBooks(c *gin.Context) {
	address := fmt.Sprintf("%v:%v", os.Getenv("BOOK_SERVICE_HOST"), os.Getenv("BOOK_SERVICE_PORT"))
	conn, err := grpc.Dial(
		address,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second,
	)
	defer cancel()

	client := proto.NewBookServiceClient(conn)
	getBooksRequest := proto.GetBooksRequest{}

	res, err := client.GetBooks(ctx, &getBooksRequest)
	if err != nil {
		log.Fatal("Request failed.")
		return
	}

	c.JSONP(http.StatusOK, gin.H{
		"books": res.Books,
	})
}
