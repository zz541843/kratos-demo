package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"kratos-demo/internal/biz"
)

type User struct {
	gorm.Model
	Name   string   `gorm:"type:varchar(100);not null"`
	Phone  string   `gorm:"type:varchar(100);not null"`
	Orders []*Order `gorm:"foreignkey:UserID"`
}

type Order struct {
	gorm.Model
	UserID  uint `gorm:"not null"`
	User    *User
	OrderNo string `gorm:"type:varchar(100);not null"`
}

type userRepo struct {
	data *Data
	log  *log.Helper
}

func (u userRepo) Save(ctx context.Context, user *biz.User) (*biz.User, error) {
	return user, nil
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
