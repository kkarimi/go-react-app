package main

import (
	r "github.com/dancannon/gorethink"
	"log"
	"net/http"
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

	// User routes
	// router.Handle("user edit", editUser)
	// router.Handle("user subscribe", subscribeUser)
	// router.Handle("user unsubscribe", unsubscribeIser)

	// Message routes
	println("Running...")
	http.Handle("/", router)
	http.ListenAndServe(BackendPort, nil)
}
