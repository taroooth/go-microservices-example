package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
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

func CreateBook(c *gin.Context) {
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

	var book proto.Book
	if err := c.BindJSON(&book); err != nil {
        return
    }

	client := proto.NewBookServiceClient(conn)
	createBookRequest := proto.CreateBookRequest{
		AuthorId: book.AuthorId,
		Title: book.Title,
	}

	res, err := client.CreateBook(ctx, &createBookRequest)
	if err != nil {
		log.Fatal("Request failed.")
		return
	}

	c.JSONP(http.StatusCreated, gin.H{
		"book": res.Book,
	})
}

func UpdateBook(c *gin.Context) {
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

	var book proto.Book
	if err := c.BindJSON(&book); err != nil {
        return
    }

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		log.Fatal("Request failed.")
		return
	}

	client := proto.NewBookServiceClient(conn)
	updateBookRequest := proto.UpdateBookRequest{
		Id: id,
		AuthorId: book.AuthorId,
		Title: book.Title,
	}

	_, err = client.UpdateBook(ctx, &updateBookRequest)
	if err != nil {
		log.Fatal("Request failed.")
		return
	}

	c.JSONP(http.StatusNoContent, gin.H{
		"message": "Book updated successfully.",
	})
}

func DeleteBook(c *gin.Context) {
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

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		log.Fatal("Request failed.")
		return
	}

	client := proto.NewBookServiceClient(conn)
	deleteBookRequest := proto.DeleteBookRequest{
		Id: id,
	}

	_, err = client.DeleteBook(ctx, &deleteBookRequest)
	if err != nil {
		log.Fatal("Request failed.")
		return
	}

	c.JSONP(http.StatusNoContent, gin.H{
		"message": "Book deleted successfully.",
	})
}
