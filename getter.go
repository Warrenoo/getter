package getter

import (
	"flag"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"time"
)

type MateData struct {
	data string
	st   time.Time
}

func (this *MateData) Data() string {
	return this.data
}

func (this *MateData) St() time.Time {
	return this.st
}

type Client struct {
	url  string
	path string
	conn *websocket.Conn
	ch   chan *MateData
	done chan bool
}

const channelBufSize = 1024

func New(url string, path string) *Client {
	ch := make(chan *MateData, channelBufSize)
	done := make(chan bool, 1)

	return &Client{url, path, nil, ch, done}
}

func (this *Client) Conn() *websocket.Conn {
	return this.conn
}

func (this *Client) Ch() chan *MateData {
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

			if err != nil {
				this.done <- true
			} else {
				md := MateData{string(message), time.Now()}

				this.ch <- &md
			}
		}
	}
}

func (this *Client) OnOpen(f func(*Client)) {
	var addr = flag.String("addr", this.url, "http service address")
	u := url.URL{Scheme: "ws", Host: *addr, Path: this.path}
	ws, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	sign_log(err)
	this.conn = ws
	f(this)
}

func (this *Client) OnListen(messages *[]string, f func(*Client)) {
	for _, v := range *messages {
		this.send([]byte(v))
	}

	go this.receive()
	f(this)
}

func (this *Client) OnClose(f func(*Client)) {
	f(this)
	this.conn.Close()
}

func sign_log(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
