package websocket

import "fmt"

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			thisReadableName := client.ReadableName
			for client, _ := range pool.Clients {
				fmt.Printf("ClientID: %s \n Name: %s\n", client.ID, client.ReadableName)
				client.Conn.WriteJSON(Message{Type: 1, Body: fmt.Sprintf("New User Joined: %s", thisReadableName)})
			}
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected..."})
			}
			break
		case message := <-pool.Broadcast:
			fmt.Println("Sending message to all clients in Pool")
			for client, _ := range pool.Clients {
				//If not checked the message sent gets resent to sending client
				//this gives a dublicate message on the senders screen
				if client.ID != message.Sender {
					client.Conn.WriteJSON(Message{Type: 1, Body: message.Body, ReadableName: message.ReadableName})
				}

			}
		}
	}
}
