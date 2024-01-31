package main

import (
	"fmt"
	"net/http"

	"github.com/Tommy-56/realtime-chat-go-react/pkg/websocket"
	"github.com/google/uuid"
	"github.com/goombaio/namegenerator"
)

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint Hit")
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}
	clientID := uuid.New()
	readableName := namegenerator.NewNameGenerator(int64(clientID.Time())).Generate()
	client := &websocket.Client{
		ID:           clientID.String(),
		Conn:         conn,
		Pool:         pool,
		ReadableName: readableName,
	}

	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

func main() {
	fmt.Println("Chat App v0.1")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
