package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "kratos-demo/api/helloworld/v1"
	"kratos-demo/internal/biz"
)

// UserService is a greeter service.
type UserService struct {
	v1.UnimplementedUserServer

	uc *biz.UserUsecase
}

// NewUserService new a greeter service.
func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{uc: uc}
}

func (u *UserService) GetUser(ctx context.Context, in *v1.GetUserRequest) (*v1.UserResponse, error) {
	user, err := u.uc.CreateUser(ctx, &biz.User{
		Name: "张三",
	})
	if err != nil {
		return nil, err
	}
	return &v1.UserResponse{
		Name: user.Name,
	}, nil
}
func (u *UserService) CreateUser(context.Context, *v1.CreateUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (u *UserService) UpdateUser(context.Context, *v1.UpdateUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (u *UserService) DeleteUser(context.Context, *v1.DeleteUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (u *UserService) ListUsers(context.Context, *v1.ListUsersRequest) (*v1.ListUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (u *UserService) CreateOrderForUser(context.Context, *v1.CreateOrderForUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrderForUser not implemented")
}
func (u *UserService) DeleteOrderForUser(context.Context, *v1.DeleteOrderForUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrderForUser not implemented")
}
