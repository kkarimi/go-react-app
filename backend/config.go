package main

import (
	"os"
)

var RethinkHost = os.Getenv("RETHINK_DB_HOST")
var BackendPort = os.Getenv("BACKEND_PORT")
