package grpc

import (
	"context"

	"github.com/taroooth/order/services/book/proto"
)

type server struct {
	proto.UnimplementedBookServiceServer
}

func NewBookService() proto.BookServiceServer {
	return &server{}
}

func (s *server) GetBooks(ctx context.Context, req *proto.GetBooksRequest) (*proto.GetBooksResponse, error) {
	res := &proto.GetBooksResponse{
		Books: make([]*proto.Book, 10),
	}

	return res, nil
}
