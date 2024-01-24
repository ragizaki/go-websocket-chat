package user

import "context"

type User struct {
	ID       int64  `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Email    string `json:"email" db:"email"`
}

type CreateUserRequest struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Email    string `json:"email" db:"email"`
}

type CreateUserResponse struct {
	ID       string `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
}

type Repository interface {
	CreateUser(context.Context, *User) (*User, error)
}

type Service interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
}
