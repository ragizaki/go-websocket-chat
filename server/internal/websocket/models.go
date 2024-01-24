package websocket

type CreateRoomRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type RoomResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ClientResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}
