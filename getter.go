package getter

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
)

type Client struct {
	url  string
	path string
	conn *websocket.Conn
	ch   chan string
	done chan bool
}

const channelBufSize = 1024

func New(url string, path string) *Client {
	ch := make(chan string, channelBufSize)
	done := make(chan bool)

	return &Client{url, path, nil, ch, done}
}

func (this *Client) Ch() chan string {
	return this.ch
}

func (this *Client) Done() chan bool {
	return this.done
}

func (this *Client) send(message []byte) {
	err := this.conn.WriteMessage(websocket.TextMessage, message)
	sign_log(err)
}

func (this *Client) receive() {
	for {
		select {
		// receive done request
		case <-this.done:
			return
		// read data from websocket connection
		default:
			_, message, err := this.conn.ReadMessage()
			sign_log(err)
			fmt.Printf("pid1: %d\n", os.Getpid())

			if err != nil {
				this.done <- true
			} else {
				this.ch <- string(message)
			}
		}
	}
}

func (this *Client) OnOpen(f func()) {
	var addr = flag.String("addr", this.url, "http service address")
	u := url.URL{Scheme: "ws", Host: *addr, Path: this.path}
	ws, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	sign_log(err)
	this.conn = ws
	f()
}

func (this *Client) OnListen(message []byte, f func()) {
	this.send(message)

	go this.receive()
	f()
}

func (this *Client) OnClose(f func()) {
	this.conn.Close()
	f()
}

func sign_log(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
