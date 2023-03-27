package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Upgrader will upgrade the connection from http to a WebSocket.
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}

// Reader will listen for new messages from the WebSocket endpoint
func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func serveWS(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	//Upgrades the connection from http to WS
	websocket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	//This will listen indefinitely for new messages arriving through the WS
	reader(websocket)
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "WEST SIDE")
	})

	http.HandleFunc("/ws", serveWS)
}
