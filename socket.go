package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server...")
	http.Handle("/", http.FileServer(http.Dir("./public/")))
	http.Handle("/ws", websocket.Handler(socketHandler))
	fmt.Println("Listening on port 8888")
	http.ListenAndServe(":8888", nil)
}

// slice for socket connections
var cs []*websocket.Conn

func socketHandler(ws *websocket.Conn) {
	// add socket connection to slice
	cs = append(cs, ws)
	var msg []byte
	// loop until connection is closed
	for {
		if err := websocket.Message.Receive(ws, &msg); err != nil {
			fmt.Println("Connection closed")
			return
		}
		// send message to each connection
		for _, c := range cs {
			websocket.Message.Send(c, string(msg))
		}
		fmt.Println("Received: " + string(msg))
	}
}
