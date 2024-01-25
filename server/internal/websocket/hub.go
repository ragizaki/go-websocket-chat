package websocket

type Room struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Owner   string             `json:"owner"`
	Clients map[string]*Client `json:"clients"`
}

type Hub struct {
	Rooms       map[string]*Room
	RoomNames   map[string]struct{}
	Subscribe   chan *Client
	Unsubscribe chan *Client
	Broadcast   chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Rooms:       make(map[string]*Room),
		RoomNames:   make(map[string]struct{}),
		Subscribe:   make(chan *Client),
		Unsubscribe: make(chan *Client),
		Broadcast:   make(chan *Message, 5),
	}
}

func (hub *Hub) CheckRoomExists(name string) bool {
	if _, exists := hub.RoomNames[name]; exists {
		return true
	}
	return false
}

func (hub *Hub) AddRoom(id string, name string, owner string) {
	room := &Room{
		ID:      id,
		Name:    name,
		Owner:   owner,
		Clients: make(map[string]*Client),
	}

	hub.Rooms[id] = room
	hub.RoomNames[name] = struct{}{}
}

func handleSubscribe(hub *Hub, client *Client) {
	if room, roomExists := hub.Rooms[client.RoomID]; roomExists {
		if _, clientExists := room.Clients[client.ID]; !clientExists {
			room.Clients[client.ID] = client
		}
	}
}

func handleUnsubscribe(hub *Hub, client *Client) {
	if room, roomExists := hub.Rooms[client.RoomID]; roomExists {
		if _, clientExists := room.Clients[client.ID]; clientExists {
			if len(room.Clients) > 0 {
				// inform other users that user has left the room
				hub.Broadcast <- &Message{
					Content:  "User " + client.Username + " has left the room",
					RoomID:   client.RoomID,
					Username: client.Username,
				}
			}
			delete(room.Clients, client.ID)
			close(client.Message)
		}
	}
}

func handleBroadcast(hub *Hub, message *Message) {
	if room, roomExists := hub.Rooms[message.RoomID]; roomExists {
		for _, client := range room.Clients {
			client.Message <- message
		}
	}
}
func (hub *Hub) Run() {
	for {
		select {
		case client := <-hub.Subscribe:
			handleSubscribe(hub, client)
		case client := <-hub.Unsubscribe:
			handleUnsubscribe(hub, client)
		case message := <-hub.Broadcast:
			handleBroadcast(hub, message)
		}
	}
}
