package main

import (
	"fmt"
	"github.com/warrenoo/getter"
)

func main() {

	var origin string = "http://test.com"
	var url string = "ws://127.0.0.1:6000/websocket"

	client := getter.New(origin, url, "")

	client.OnConnect(func() {
		fmt.Printf("Init: %s\n", url)
	})

	message := []byte("{'test':1}")
	fmt.Printf("Listen: %s\n", message)
	client.OnListen(message, func(data []byte) {
		fmt.Printf("Receive: %s\n", data)
	})

	client.OnClose(func() {
		fmt.Printf("Close Server!!")
	})
}
