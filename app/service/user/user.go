package user

import (
	"context"

	"github.com/taroooth/order/app/pb"
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
