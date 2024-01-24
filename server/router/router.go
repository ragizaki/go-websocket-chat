package router

import (
	"server/internal/user"

	"github.com/gin-gonic/gin"
)

func NewRouter(userHandler *user.Handler) *gin.Engine {
	router := gin.Default()

	router.POST("/signup", userHandler.CreateUser)

	return router
}

func StartRouter(addr string, router *gin.Engine) error {
	return router.Run(addr)
}
