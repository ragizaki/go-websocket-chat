package user

import (
	"context"
	"server/config"
	"server/util"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const HOURS_IN_DAY = 24

type service struct {
	Repository
	timeout time.Duration
}

type JWTClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
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

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	userRes, err := s.Repository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	res := &CreateUserResponse{
		ID:       strconv.Itoa(int(userRes.ID)),
		Username: userRes.Username,
		Email:    userRes.Email,
	}

	return res, nil
}

func (s *service) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	user, err := s.Repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	err = util.ComparePassword(user.Password, req.Password)
	if err != nil {
		return nil, err
	}

	userID := strconv.Itoa(int(user.ID))
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256, JWTClaims{
			ID:       userID,
			Username: user.Username,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    userID,
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(HOURS_IN_DAY * time.Hour)),
			},
		},
	)
	secret := config.GetEnv("JWT_SECRET")
	signedString, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	res := &LoginResponse{
		accessToken: signedString,
		ID:          userID,
		Username:    user.Username,
	}

	return res, nil
}
