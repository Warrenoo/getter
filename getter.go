package getter

import (
	"golang.org/x/net/websocket"
	"log"
)

type Data []byte

type Client struct {
	origin string
	url    string
	ext    string
	conn   *websocket.Conn
	ch     chan Data
	done   chan bool
}

const channelBufSize = 1024

func New(origin string, url string, ext string) *Client {
	ch := make(chan Data, channelBufSize)
	done := make(chan bool)

	return &Client{origin, url, ext, nil, ch, done}
}

func (this *Client) Ch() chan Data {
	return this.ch
}

func (this *Client) Done() chan bool {
	return this.done
}

func (this *Client) send(message []byte) {
	_, err := this.conn.Write(message)
	sign_log(err)
}

func (this *Client) receive() {
	var data Data = make(Data, 2048)

	for {
		select {
		// receive done request
		case <-this.done:
			this.done <- true
			return
		// read data from websocket connection
		default:

			m, err := this.conn.Read(data)
			sign_log(err)

			if err != nil {
				this.done <- true
			} else {
				this.ch <- data[:m]
			}
		}
	}
}

func (this *Client) OnConnect(f func()) {
	ws, err := websocket.Dial(this.url, this.ext, this.origin)
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
