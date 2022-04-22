package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"kratos-demo/internal"

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

func (u *UserService) GetUser(ctx context.Context, in *v1.GetUserRequest) (out *v1.UserResponse, err error) {
	out = &v1.UserResponse{}
	user, err := u.uc.GetUserById(ctx, in.Id)
	err = internal.Copier.StructCopy(out, *user)
	return out, errors.New(300, "321", "213")
}

func (u *UserService) CreateUser(ctx context.Context, in *v1.CreateUserRequest) (*emptypb.Empty, error) {
	bizIn := biz.User{}

	err := internal.Copier.StructCopy(&bizIn, in)
	if err != nil {
		return nil, err
	}
	err = u.uc.CreateUser(ctx, &bizIn)
	if err != nil {
		return nil, err
	}
	return nil, err
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
