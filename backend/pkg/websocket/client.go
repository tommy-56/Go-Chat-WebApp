package websocket

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID           string
	Conn         *websocket.Conn
	Pool         *Pool
	ReadableName string
}

type Message struct {
	Type         int    `json:"type"`
	Body         string `json:"body"`
	Sender       string `json:"sender"`
	ReadableName string `json:"readableName"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		message := Message{Type: messageType, Body: string(p), Sender: c.ID, ReadableName: c.ReadableName}

		c.Pool.Broadcast <- message

		fmt.Printf("Message Received: %+v\n", message)
		err = c.Conn.WriteJSON(message)
		if err != nil {
			fmt.Println(err)
		}

	}
}
