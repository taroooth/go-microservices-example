package grpc

import (
	"context"

	"github.com/taroooth/order/services/book/model"

	"github.com/taroooth/order/services/book/proto"
)

var _ proto.BookServiceServer = (*server)(nil)

type server struct {
	proto.UnimplementedBookServiceServer
}

func NewBookService() proto.BookServiceServer {
	return &server{}
}

func (s *server) GetBooks(ctx context.Context, req *proto.GetBooksRequest) (*proto.GetBooksResponse, error) {
	var books []model.Book
	model.GetAllBooks(&books)

	res := &proto.GetBooksResponse{
		Books: make([]*proto.Book, len(books)),
	}

	for i, book := range books {
		res.Books[i] = &proto.Book{
			Id:       book.Id,
			AuthorId: book.AuthorId,
			Title:    book.Title,
		}
	}

	return res, nil
}

func (s *server) GetBook(ctx context.Context, req *proto.GetBookRequest) (*proto.GetBookResponse, error) {
	return &proto.GetBookResponse{
		Book: &proto.Book{
			Id:       2,
			AuthorId: 2,
			Title:    "Title2",
		},
	}, nil
}
