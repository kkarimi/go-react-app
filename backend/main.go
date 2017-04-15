package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

type Channel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

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

	for {
		var inMessage Message
		if err := socket.ReadJSON(&inMessage); err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("%#v\n", inMessage)
	}
}
