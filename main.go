package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/tamizhvendan/gomidway/post"
	"github.com/tamizhvendan/gomidway/post/publish"
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
	publishPost(db)
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
	fmt.Printf("Ok: User '%s' logged in\n", res.User.Username)
}

func publishPost(db *gorm.DB) {
	res, err := publish.NewPost(db, &publish.Request{
		AuthorId: 1,
		Body:     "Golang rocks!",
		Title:    "My first gomidway post",
		Tags:     []string{"intro", "golang"},
	})
	if err != nil {
		if _, ok := err.(*post.TitleDuplicateError); ok {
			fmt.Println("Bad Request: ", err.Error())
			return
		}
		fmt.Println("Internal Server Error: ", err.Error())
		return
	}
	fmt.Println("Created : ", res.PostId)
}
