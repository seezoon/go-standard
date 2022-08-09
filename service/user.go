package service

import (
	"context"
	"go-standard/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type userService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() pb.UserServiceServer {
	return &userService{}
}

func (s *userService) AddUser(ctx context.Context, request *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *userService) ModifyUserMobile(ctx context.Context, request *pb.ModifyMobileRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (s *userService) QryUserById(ctx context.Context, request *pb.QryUserByIdRequest) (*pb.UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *userService) QryUserPage(ctx context.Context, request *pb.QryUserPageRequest) (*pb.QryUserPageResponse, error) {
	//TODO implement me
	panic("implement me")
}
