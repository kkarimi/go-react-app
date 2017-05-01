package main

import (
	r "github.com/dancannon/gorethink"
	"log"
	"net/http"
)

//Channel struct
type Channel struct {
	ID   string `json:"id" gorethink:"id,omitempty"`
	Name string `json:"name" gorethink:"name"`
}

type User struct {
	ID   string `gorethink:"id,omitempty"`
	Name string `gorethink:"name"`
}

func main() {
	session, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "chatserver",
	})
	if err != nil {
		log.Panic(err.Error())
	}
	router := NewRouter(session)
	router.Handle("channel add", addChannel)
	http.Handle("/", router)
	http.ListenAndServe(":4000", nil)
}
