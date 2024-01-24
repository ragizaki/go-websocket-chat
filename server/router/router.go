package router

import (
	"server/internal/user"
	"server/internal/websocket"

	"github.com/gin-gonic/gin"
)

func NewRouter(userHandler *user.Handler, websocketHandler *websocket.Handler) *gin.Engine {
	router := gin.Default()

	// user routes
	router.POST("/signup", userHandler.CreateUser)
	router.POST("/login", userHandler.Login)
	router.GET("/logout", userHandler.Logout)

	// websocket routes
	router.POST("/ws/rooms", websocketHandler.CreateRoom)
	router.GET("/ws/rooms/join/:roomId", websocketHandler.JoinRoom)
	router.GET("/ws/rooms", websocketHandler.GetRooms)
	router.GET("/ws/clients/:roomId", websocketHandler.GetClientsInRoom)

	return router
}

func StartRouter(addr string, router *gin.Engine) error {
	return router.Run(addr)
}
