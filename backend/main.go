package main

import (
	"log"
	"net/http"

	r "github.com/dancannon/gorethink"
)

func main() {

	if RethinkHost != "" {
		RethinkHost = "db:28015"
	}

	if BackendPort != "" {
		BackendPort = ":4000"
	}

	session, err := r.Connect(r.ConnectOpts{
		Address:  RethinkHost,
		Database: "chatserver",
	})
	if err != nil {
		log.Panic(err.Error())
	}
	router := NewRouter(session)

	router.Handle("channel add", addChannel)
	router.Handle("channel subscribe", subscribeChannel)
	router.Handle("channel unsubscribe", unsubscribeChannel)

	router.Handle("user edit", editUser)
	router.Handle("user subscribe", subscribeUser)
	router.Handle("user unsubscribe", unsubscribeUser)

	router.Handle("message add", addChannelMessage)
	router.Handle("message subscribe", subscribeChannelMessage)
	router.Handle("message unsubscribe", unsubscribeChannelMessage)

	// Message routes
	println("Running...")
	http.Handle("/", router)
	http.ListenAndServe(BackendPort, nil)
}
