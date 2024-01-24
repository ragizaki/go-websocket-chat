package user

import (
	"net/http"
	"server/config"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type Handler struct {
	Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service,
	}
}

func (handler *Handler) CreateUser(ctx *gin.Context) {
	var userReq CreateUserRequest
	if err := ctx.ShouldBindJSON(&userReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userRes, err := handler.Service.CreateUser(ctx.Request.Context(), &userReq)
	if err != nil {
		pqErr, ok := err.(*pq.Error)
		if ok && pqErr.Code == "23505" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, userRes)
}

func (handler *Handler) Login(ctx *gin.Context) {
	var loginReq LoginRequest
	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loginRes, err := handler.Service.Login(ctx.Request.Context(), &loginReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	maxAge, _ := strconv.Atoi(config.GetEnv("JWT_MAX_AGE"))
	domain := config.GetEnv("JWT_DOMAIN")

	ctx.SetCookie("jwt", loginRes.accessToken, maxAge, "/", domain, false, true)

	loginRes = &LoginResponse{
		ID:       loginRes.ID,
		Username: loginRes.Username,
	}

	ctx.JSON(http.StatusOK, loginRes)
}
