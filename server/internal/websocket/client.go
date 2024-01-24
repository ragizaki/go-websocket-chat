package websocket

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn     *websocket.Conn
	Message  chan *Message
	ID       string `json:"id"`
	RoomID   string `json:"roomId"`
	Username string `json:"username"`
}

type Message struct {
	Content  string `json:"content"`
	RoomID   string `json:"roomId"`
	Username string `json:"username"`
}

func (client *Client) writeMessage() {
	defer func() {
		client.Conn.Close()
	}()

	for {
		message, ok := <-client.Message
		if !ok {
			return
		}
		client.Conn.WriteJSON(message)
	}
}

func (client *Client) readMessage(hub *Hub) {
	defer func() {
		hub.Unsubscribe <- client
		client.Conn.Close()
	}()

	for {
		_, message, err := client.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		msg := &Message{
			Content:  string(message),
			RoomID:   client.RoomID,
			Username: client.Username,
		}

		hub.Broadcast <- msg
	}
}
