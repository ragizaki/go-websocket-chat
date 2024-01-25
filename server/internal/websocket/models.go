package websocket

type CreateRoomRequest struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

type RoomResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

type ClientResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}
