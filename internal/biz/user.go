package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Order struct {
	Id      uint
	UserID  uint
	User    *User
	OrderNo string
}
type User struct {
	Id     uint
	Name   string
	Phone  string
	Orders []*Order
}
type UserRepo interface {
	CreateUser(context.Context, *User) error
	GerUserById(context.Context, uint32) (*User, error)
	DelUserById(context.Context, uint32) error
}

// UserUsecase is a Greeter usecase.
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// NewUserUsecase new a NewUserUsecase usecase.
func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateUser creates a Greeter, and returns the new Greeter.
func (uc *UserUsecase) CreateUser(ctx context.Context, g *User) (err error) {
	return uc.repo.CreateUser(ctx, g)
}

func (uc *UserUsecase) GetUserById(ctx context.Context, id uint32) (g *User, err error) {
	return uc.repo.GerUserById(ctx, id)
}
