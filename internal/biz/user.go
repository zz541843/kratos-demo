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
	Save(context.Context, *User) (*User, error)
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
func (uc *UserUsecase) CreateUser(ctx context.Context, g *User) (*User, error) {
	uc.log.WithContext(ctx).Infof("Createuser: %v", g.Name)
	return uc.repo.Save(ctx, g)
}
