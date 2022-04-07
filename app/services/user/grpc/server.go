package grpc

import (
	"context"

	"github.com/taroooth/order/app/services/user/pb"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() pb.UserServiceServer {
	return &server{}
}

func (s *server) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	return &pb.CreateResponse{
		Token: "aaaaaaaaaaaaaa",
	}, nil
}
