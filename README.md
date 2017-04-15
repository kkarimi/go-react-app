# go-react-app
Playing around with Go and React

To run:
1- `$ cd backend; go run main.go`

Packages needed (install with `go get`):

1- [gorilla/websocket](https://github.com/gorilla/websocket) ([docs](https://godoc.org/github.com/gorilla/websocket))
2- [mitchellh/mapstructure](https://github.com/mitchellh/mapstructure) ([docs](https://godoc.org/github.com/mitchellh/mapstructure))

Testing
--
Send test JSON to websocket by running `node frontend/src/testsocket.js` 
Can also test websockets with the [echo test](https://websocket.org/echo.html)
