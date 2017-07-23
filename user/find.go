package user

import "github.com/jinzhu/gorm"

type EmailNotExistsError struct{}

func (*EmailNotExistsError) Error() string {
	return "email not exists"
}

func FindByEmail(db *gorm.DB, email string) (*User, error) {
	var user User
	res := db.Find(&user, &User{Email: email})
	if res.RecordNotFound() {
		return nil, &EmailNotExistsError{}
	}
	return &user, nil
}
