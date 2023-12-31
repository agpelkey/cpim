package main

import (
	"fmt"
	"log"

	"golang.org/x/net/websocket"
)


func main() {

	origin := "http://localhost/"
	url := "ws://localhost:8080/echo"

	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	// write bytes to the server
	if _, err := ws.Write([]byte("Hello World\n")); err != nil {
		log.Fatal(err)
	}

	var msg = make([]byte, 512)
	var n int
	if n, err = ws.Read(msg); err != nil {
		log.Fatal(err)
	}

	// print the message string from beginning to 'n' number of bytes
	fmt.Printf("Received: %s\n", msg[:n])
}
