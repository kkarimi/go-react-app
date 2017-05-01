package main

import (
	"fmt"
	r "github.com/dancannon/gorethink"
	"github.com/gorilla/websocket"
	"net/http"
)

// Handler interface empty payload
// if error, need a client to send err message back
type Handler func(*Client, interface{})

// handle cross-origin by func
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

//Router struct
type Router struct {
	rules   map[string]Handler
	session *r.Session
}

//NewRouter func
func NewRouter(session *r.Session) *Router {
	return &Router{
		rules:   make(map[string]Handler),
		session: session,
	}
}

//Handle func
func (r *Router) Handle(msgName string, handler Handler) {
	// take Router pointer which attaches this to the struct
	r.rules[msgName] = handler
}

//FindHandler func
func (r *Router) FindHandler(msgName string) (Handler, bool) {
	// look up handler from rules
	handler, found := r.rules[msgName]
	return handler, found
}

//ServeHTTP func
func (r *Router) ServeHTTP(w http.ResponseWriter, q *http.Request) {
	// upgrade http to websocket
	// hijack connection and switch protocol from http to websockets
	socket, err := upgrader.Upgrade(w, q, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
	client := NewClient(socket, r.FindHandler, r.session)
	defer client.Close()
	go client.Write()
	client.Read()
}
