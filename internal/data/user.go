package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"kratos-demo/internal"
	"kratos-demo/internal/biz"
	"reflect"
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

func (u userRepo) GerUserById(ctx context.Context, u2 uint32) (out *biz.User, err error) {
	out = &biz.User{}
	if reflect.TypeOf(out).Kind() != reflect.Ptr || reflect.TypeOf(out).Elem().Kind() != reflect.Struct {
		fmt.Println(reflect.TypeOf(out).Kind())
	}
	user := &User{}
	u.data.db.First(user, u2)
	err = internal.Copier.StructCopy(out, *user)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return
}

func (u userRepo) DelUserById(ctx context.Context, u2 uint32) error {
	//TODO implement me
	panic("implement me")
}

// CreateUser 创建用户 by userRepo
func (u userRepo) CreateUser(ctx context.Context, in *biz.User) (err error) {
	user := User{}
	err = internal.Copier.StructCopy(&user, in)
	if err != nil {
		return err
	}
	u.data.db.WithContext(ctx).Create(&user)
	return nil
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
