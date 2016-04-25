package getter

import (
	"github.com/golang/net/websocket"
	"log"
)

type Client struct {
	origin string
	url    string
	ext    string
	conn   *websocket.Conn
}

func New(origin string, url string, ext string) *Client {
	client := &Client{origin, url, ext, nil}

	//runtime.SetFinalizer(client, client.conn.Close)
	return client
}

func (this *Client) send(message []byte) {
	_, err := this.conn.Write(message)
	sign_log(err)
}

func (this *Client) receive() []byte {
	var data []byte = make([]byte, 2048)

	m, err := this.conn.Read(data)
	sign_log(err)
	return data[:m]
}

func (this *Client) OnConnect(f func()) {
	ws, err := websocket.Dial(this.url, this.ext, this.origin)
	sign_log(err)
	this.conn = ws
	f()
}

func (this *Client) OnListen(message []byte, f func([]byte)) {
	this.send(message)

	for true {
		data := this.receive()
		f(data)
	}

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
