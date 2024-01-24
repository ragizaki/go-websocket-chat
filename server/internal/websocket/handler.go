package websocket

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct {
	hub *Hub
}

func NewHandler(hub *Hub) *Handler {
	return &Handler{
		hub: hub,
	}
}

func (handler *Handler) CreateRoom(ctx *gin.Context) {
	var req CreateRoomRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	handler.hub.AddRoom(req.ID, req.Name)

	ctx.JSON(http.StatusOK, req)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(req *http.Request) bool {
		return true
	},
}

func (handler *Handler) JoinRoom(ctx *gin.Context) {
	connection, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	roomId := ctx.Param("roomId")
	clientId := ctx.Param("userId")
	username := ctx.Param("username")

	client := &Client{
		Conn:     connection,
		Message:  make(chan *Message, 10),
		ID:       clientId,
		RoomID:   roomId,
		Username: username,
	}

	// inform other users that new user has joined a room
	message := &Message{
		Content:  "User " + username + " has joined the room",
		RoomID:   roomId,
		Username: username,
	}

	// adds the client to the room
	handler.hub.Subscribe <- client

	// broadcasts the message to all clients in the room
	handler.hub.Broadcast <- message

	go client.writeMessage()
	client.readMessage(handler.hub)
}

func (handler *Handler) GetRooms(ctx *gin.Context) {
	rooms := make([]RoomResponse, 0)

	for _, room := range handler.hub.Rooms {
		rooms = append(rooms, RoomResponse{
			ID:   room.ID,
			Name: room.Name,
		})
	}
	ctx.JSON(http.StatusOK, rooms)
}

func (handler *Handler) GetClientsInRoom(ctx *gin.Context) {
	clients := make([]ClientResponse, 0)
	roomId := ctx.Param("roomId")

	if _, roomExists := handler.hub.Rooms[roomId]; !roomExists {
		ctx.JSON(http.StatusOK, clients)
		return
	}
	for _, client := range handler.hub.Rooms[roomId].Clients {
		clients = append(clients, ClientResponse{
			ID:       client.ID,
			Username: client.Username,
		})
	}
	ctx.JSON(http.StatusOK, clients)
}
