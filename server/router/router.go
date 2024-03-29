package router

import (
	"server/internal/user"
	"server/internal/websocket"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(userHandler *user.Handler, websocketHandler *websocket.Handler) *gin.Engine {
	router := gin.Default()

	// CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// define prefixes for routes (/api/auth and /api/ws)
	api := router.Group("/api")

	// user routes
	authRouter := api.Group("/auth")
	authRouter.POST("/signup", userHandler.CreateUser)
	authRouter.POST("/login", userHandler.Login)
	authRouter.GET("/logout", userHandler.Logout)

	// websocket routes
	wsRouter := api.Group("/ws")
	wsRouter.POST("/rooms", websocketHandler.CreateRoom)
	wsRouter.GET("/rooms", websocketHandler.GetRooms)
	router.GET("/api/ws/rooms/:roomId/join", websocketHandler.JoinRoom)
	wsRouter.GET("/rooms/:roomId/clients", websocketHandler.GetClientsInRoom)

	return router
}

func StartRouter(addr string, router *gin.Engine) error {
	return router.Run(addr)
}
