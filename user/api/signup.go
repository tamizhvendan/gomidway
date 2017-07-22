package api

import (
	"github.com/jinzhu/gorm"
	"github.com/tamizhvendan/gomidway/user"
	"golang.org/x/crypto/bcrypt"
)

type SignupUser struct {
	Username string
	Email    string
	Password string
}

type UserSignedup struct {
	Id uint
}

func Signup(db *gorm.DB, req *SignupUser) (*UserSignedup, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	newUser := &user.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: string(passwordHash),
	}

	id, err := user.Create(db, newUser)
	if err != nil {
		return nil, err
	}
	return &UserSignedup{Id: id}, err
}
