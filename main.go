package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/tamizhvendan/gomidway/user"
	"github.com/tamizhvendan/gomidway/user/login"
	"github.com/tamizhvendan/gomidway/user/signup"
)

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := gorm.Open("postgres",
		`host=localhost 
			user=postgres password=test
			dbname=gomidway 
			sslmode=disable`)
	panicOnError(err)
	defer db.Close()

	signupUser(db)
	loginUser(db)

}

func signupUser(db *gorm.DB) {
	res, err := signup.Signup(db, &signup.Request{
		Email:    "foo@bar.com",
		Username: "foo",
		Password: "foobar",
	})
	if err != nil {
		switch err.(type) {
		case *user.UsernameDuplicateError:
			fmt.Println("Bad Request: ", err.Error())
			return
		case *user.EmailDuplicateError:
			fmt.Println("Bad Request: ", err.Error())
			return
		default:
			fmt.Println("Internal Server Error: ", err.Error())
			return
		}
	}
	fmt.Println("Created: ", res.Id)
}

func loginUser(db *gorm.DB) {
	res, err := login.Login(db, &login.Request{Email: "foo@bar.com", Password: "foobar"})
	if err != nil {
		switch err.(type) {
		case *user.EmailNotExistsError:
			fmt.Println("Bad Request: ", err.Error())
			return
		case *login.PasswordMismatchError:
			fmt.Println("Bad Request: ", err.Error())
			return
		default:
			fmt.Println("Internal Server Error: ", err.Error())
			return
		}
	}
	fmt.Printf("Ok: User '%s' logged in", res.User.Username)
}

/*
Created:  1
Bad Request:  Username 'foo' already exists
Bad Request:  Email 'foo@bar.com' already exists
Created:  4
Ok: User 'foo' logged in
Bad Request:  password didn't match
*/
