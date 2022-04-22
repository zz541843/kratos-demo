package main

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"testing"

	gerrors "errors"
)

type Option func(*options)

type options struct {
	Name string
}

type CustomErr struct {
	Message string
	err     error
}

func (c *CustomErr) Error() string {
	return "custom error:" + c.Message
}
func (c *CustomErr) Unwrap() error {
	return c.err
}
func (c *CustomErr) Is(target error) bool {
	_, ok := target.(*CustomErr)
	if !ok {
		return false
	}
	return target.(*CustomErr).Message == c.Message
}
func TestName(t *testing.T) {
	c := CustomErr{
		Message: "test",
		err:     fmt.Errorf("asdf"),
	}
	err := gerrors.Unwrap(&c)
	fmt.Println(err)
}
func Test2(t *testing.T) {
	var e error
	var b Error
	b = "1"
	e = fmt.Errorf("asdf:%w", &CustomErr{
		Message: "123",
		err:     b,
	})
	if gerrors.As(e, &b) {
		fmt.Println(e.(*Error).Error())
		fmt.Println(e.Error())
	} else {
		fmt.Println(213)
	}
	/*var b Error
	b = "1"
	fmt.Println(gerrors.As(e, &b))*/
}

type Error string

func (e Error) Error() string { return string(e) }

type HTTPError struct {
	Errors map[string][]string `json:"errors"`

	Code int `json:"-"`
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTPError: %d", e.Code)
}
func Test4(t *testing.T) {
	FromError(&HTTPError{
		Errors: map[string][]string{
			"name": {"name is required"},
		},
		Code: 401,
	})
}
func FromError(err error) *HTTPError {
	if err == nil {
		return nil
	}
	se := &options{}
	b := &se
	c := &b
	fmt.Println(c)
	if errors.As(err, &se) {
		fmt.Println(se)
	}
	return nil
}
