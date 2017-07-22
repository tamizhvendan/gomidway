package user

import "fmt"

const (
	UniqueConstraintUsername = "users_username_key"
	UniqueConstraintEmail    = "users_email_key"
)

type User struct {
	ID           uint `gorm:"primary_key"`
	Username     string
	Email        string
	PasswordHash string
}

type UsernameDuplicateError struct {
	Username string
}

func (e *UsernameDuplicateError) Error() string {
	return fmt.Sprintf("Username '%s' already exists", e.Username)
}

type EmailDuplicateError struct {
	Email string
}

func (e *EmailDuplicateError) Error() string {
	return fmt.Sprintf("Email '%s' already exists", e.Email)
}

/*
gomidway=# \d users
                                    Table "public.users"
    Column     |          Type          |                     Modifiers
---------------+------------------------+----------------------------------------------------
 id            | integer                | not null default nextval('users_id_seq'::regclass)
 username      | character varying(50)  | not null
 email         | character varying(255) | not null
 password_hash | text                   | not null
Indexes:
    "users_pkey" PRIMARY KEY, btree (id)
    "users_email_key" UNIQUE CONSTRAINT, btree (email)
    "users_username_key" UNIQUE CONSTRAINT, btree (username)
*/
