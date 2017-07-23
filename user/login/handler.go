package login

import (
	"github.com/jinzhu/gorm"
	"github.com/tamizhvendan/gomidway/user"
	"golang.org/x/crypto/bcrypt"
)

type Request struct {
	Email    string
	Password string
}

type Response struct {
	User *user.User
}

type PasswordMismatchError struct{}

func (e *PasswordMismatchError) Error() string {
	return "password didn't match"
}

func Login(db *gorm.DB, req *Request) (*Response, error) {
	user, err := user.FindByEmail(db, req.Email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, &PasswordMismatchError{}
	}
	return &Response{User: user}, nil
}
