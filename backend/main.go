package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"time"
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
		switch inMessage.Name {
		case "channel add":
			err := addChannel(inMessage.Data)
			if err != nil {
				outMessage := Message{"error", err}
				if err := socket.WriteJSON(outMessage); err != nil {
					fmt.Println(err)
					break
				}
			}
		case "channel subscribe":
			// use go routines to create a light thread
			// so not to block app from other actions
			go subscribeChannel(socket)
		}
	}
}

func addChannel(data interface{}) error {
	var channel Channel
	err := mapstructure.Decode(data, &channel)
	if err != nil {
		return err
	}
	channel.Id = "1"
	// fmt.Printf("%#v\n", channel)
	fmt.Println("added channel")
	return nil
}

func subscribeChannel(socket *websocket.Conn) {
	for {
		time.Sleep(time.Second * 1)
		message := Message{"channel add",
			Channel{"1", "Some new channel"}}
		socket.WriteJSON(message)
		fmt.Println("send new channel")
	}
}
