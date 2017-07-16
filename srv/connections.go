package srv

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/jmataya/hermes/models"
)

// TODO: Move this into a better structure.
var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan models.Message)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()

	clients[ws] = true

	for {
		var msg models.Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}

		broadcast <- msg
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		log.Printf("New message received")
		log.Printf("%+v", msg)
	}
}
