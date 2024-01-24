package user

import (
	"net/http"

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
