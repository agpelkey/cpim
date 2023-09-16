package main

import (
	"io"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)


func main() {

	http.Handle("/echo", websocket.Handler(EchoServer))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}

}

// Echo the data received on the websocket
func EchoServer(ws *websocket.Conn) {
	io.Copy(ws, ws)
	log.Println("Received message from: ", ws.LocalAddr())
}



// start server

// acceptLoop


// readLoop
