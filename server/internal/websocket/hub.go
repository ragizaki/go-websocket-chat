package websocket

type Room struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

type Hub struct {
	Rooms       map[string]*Room
	Subscribe   chan *Client
	Unsubscribe chan *Client
	Broadcast   chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Rooms:       make(map[string]*Room),
		Subscribe:   make(chan *Client),
		Unsubscribe: make(chan *Client),
		Broadcast:   make(chan *Message, 5),
	}
}

func (hub *Hub) AddRoom(id string, name string) {
	room := &Room{
		ID:      id,
		Name:    name,
		Clients: make(map[string]*Client),
	}

	hub.Rooms[id] = room
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
func (h *Hub) Run() {
	for {
		select {
		case cl := <-h.Subscribe:
			if _, ok := h.Rooms[cl.RoomID]; ok {
				r := h.Rooms[cl.RoomID]

				if _, ok := r.Clients[cl.ID]; !ok {
					r.Clients[cl.ID] = cl
				}
			}
		case cl := <-h.Unsubscribe:
			if _, ok := h.Rooms[cl.RoomID]; ok {
				if _, ok := h.Rooms[cl.RoomID].Clients[cl.ID]; ok {
					if len(h.Rooms[cl.RoomID].Clients) != 0 {
						h.Broadcast <- &Message{
							Content:  "user left the chat",
							RoomID:   cl.RoomID,
							Username: cl.Username,
						}
					}

					delete(h.Rooms[cl.RoomID].Clients, cl.ID)
					close(cl.Message)
				}
			}

		case m := <-h.Broadcast:
			if _, ok := h.Rooms[m.RoomID]; ok {

				for _, cl := range h.Rooms[m.RoomID].Clients {
					cl.Message <- m
				}
			}
		}
	}
}
