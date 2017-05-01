package main

import (
	"fmt"
	r "github.com/dancannon/gorethink"
	"github.com/gorilla/websocket"
)

// Do not communicate by sharing memory,
// Share memory by communicating

// send messages safely from one go routines to another

//FindHandler func
type FindHandler func(string) (Handler, bool)

//Message Struct
type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

//Client struct
type Client struct {
	send         chan Message
	socket       *websocket.Conn
	findHandler  FindHandler
	session      *r.Session
	stopChannels map[int]chan bool
}

// NewStopChannel stops new channels
func (client *Client) NewStopChannel(stopKey int) chan bool {
	client.stopForKey(stopKey)
	stop := make(chan bool)
	client.stopChannels[stopKey] = stop
	return stop
}

func (client *Client) stopForKey(key int) {
	if ch, found := client.stopChannels[key]; found {
		ch <- true
		delete(client.stopChannels, key)
	}
}

// Read Reads
func (client *Client) Read() {
	var message Message
	for {
		if err := client.socket.ReadJSON(&message); err != nil {
			break
		}
		if handler, found := client.findHandler(message.Name); found {
			handler(client, message.Data)
		}
	}
	client.socket.Close()
}

func (client *Client) Write() {
	// Reciever in paranthesis makes this a method of
	// the client
	for msg := range client.send {
		//TODO: socket.sendJSON(msg)
		fmt.Printf("%#v\n", msg)
		if err := client.socket.WriteJSON(msg); err != nil {
			break
		}
	}
	client.socket.Close()
}

// Close closes channels we don't need
func (client *Client) Close() {
	for _, ch := range client.stopChannels {
		ch <- true
	}
	close(client.send)
}

// NewClient func
// convension to initialise complex objects
// as no Constructor in Go
func NewClient(socket *websocket.Conn, findHandler FindHandler, session *r.Session) *Client {
	return &Client{
		send:         make(chan Message),
		socket:       socket,
		findHandler:  findHandler,
		session:      session,
		stopChannels: make(map[int]chan bool),
	}
}
