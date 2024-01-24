package user

import (
	"context"
	"strconv"
	"time"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewUserService(repository Repository) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}

func (s *service) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)

	defer cancel()

	// TODO: hash password
	hashedPassword := req.Password

	user := &User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	user, err := s.Repository.CreateUser(ctx, user)

	if err != nil {
		return nil, err
	}

	res := &CreateUserResponse{
		ID:       strconv.Itoa(int(user.ID)),
		Username: user.Username,
		Email:    user.Email,
	}

	return res, nil
}
