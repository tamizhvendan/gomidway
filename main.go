package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/tamizhvendan/gomidway/user"
	userApi "github.com/tamizhvendan/gomidway/user/api"
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

}

func signupUser(db *gorm.DB) {
	res, err := userApi.Signup(db, &userApi.SignupUser{
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
	fmt.Println("Ok: ", res.Id)
}

/*
Ok:  1
Bad Request:  Username 'foo' already exists
Bad Request:  Email 'foo@bar.com' already exists
Ok:  4
*/
