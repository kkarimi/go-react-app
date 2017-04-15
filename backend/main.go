package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

// handle cross-origin by func
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":4000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// hijack connection and switch protocol from http to websockets
	socket, err := upgrader.Upgrade(w, r, nil)
	// handle error from ?
	if err != nil {
		fmt.Println(err)
		return
	}
	// echo incoming message back, handle any errors
	for {
		msgType, msg, err := socket.ReadMessage()
		fmt.Println(string(msg))
		if err = socket.WriteMessage(msgType, msg); err != nil {
			fmt.Println(err)
			return
		}
	}
}
