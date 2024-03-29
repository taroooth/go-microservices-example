package grpc

import (
	"context"

	"github.com/taroooth/go-microservices-example/services/book/model"

	"github.com/taroooth/go-microservices-example/services/book/proto"
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

func (s *server) CreateBook(ctx context.Context, req *proto.CreateBookRequest) (*proto.CreateBookResponse, error) {
	book := &model.Book{
		AuthorId: req.AuthorId,
		Title:    req.Title,
	}
	model.CreateBook(book)

	res := &proto.CreateBookResponse{
		Book: &proto.Book{
			Id:       book.Id,
			AuthorId: book.AuthorId,
			Title:    book.Title,
		},
	}

	return res, nil
}

func (s *server) UpdateBook(ctx context.Context, req *proto.UpdateBookRequest) (*proto.UpdateBookResponse, error) {
	book := &model.Book{
		Id:       req.Id,
		AuthorId: req.AuthorId,
		Title:    req.Title,
	}
	model.UpdateBook(book)

	res := &proto.UpdateBookResponse{
		Book: &proto.Book{
			Id:       book.Id,
			AuthorId: book.AuthorId,
			Title:    book.Title,
		},
	}

	return res, nil
}

func (s *server) DeleteBook(ctx context.Context, req *proto.DeleteBookRequest) (*proto.DeleteBookResponse, error) {
	book := &model.Book{
		Id:       req.Id,
	}
	model.DeleteBook(book)

	res := &proto.DeleteBookResponse{}

	return res, nil
}
