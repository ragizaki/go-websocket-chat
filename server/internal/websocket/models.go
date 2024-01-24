package websocket

type CreateRoomRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
