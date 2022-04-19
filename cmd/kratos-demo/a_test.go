package main

import (
	"fmt"
	jz "github.com/zz541843/go-utils"
	"testing"
)

type Option func(*options)

type options struct {
	Name string
}

func TestName(t *testing.T) {
	var a Option
	var b Option

	o := New(a, b, func(o *options) {
		o.Name = "李四"
	})
	fmt.Println(o.Name)
}
func New(opts ...Option) *options {
	o := options{
		Name: "张三",
	}
	for _, opt := range opts {
		opt(&o)
	}
	return &o
}

type A struct {
	Name string
	DD   D
}
type B struct {
	Name string
	DD   D
}
type C struct {
}
type D struct {
}

func Test3(t *testing.T) {
	a := A{
		Name: "张三",
	}
	b := B{
		Name: "李四",
	}
	newCopy := jz.NewCopy()
	err := newCopy.StructCopy(&a, b)
	if err != nil {
		return
	}

}
