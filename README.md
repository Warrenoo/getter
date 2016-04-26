# getter
- Sample WebSocket Client
## INSTALL
``````
  go get github.com/warrenoo/getter
  go get github.com/golang/net
  ln -s /usr/local/go_path/src/github.com/golang/net /usr/local/go_path/src/golang.org/x/net
  go install golang.org/x/net/websocket
``````

## USE

``````ruby
	var origin string = "http://test.com"
	var url string = "ws://127.0.0.1:6000/websocket"

	client := getter.New(origin, url, "")

	client.OnConnect(func() {
		fmt.Printf("Init: %s\n", url)
	})

	message := getter.Data("{'test':1}")
	fmt.Printf("Listen: %s\n", message)
	client.OnListen(message, func() {
		for {
			select {
			// 处理返回结果
			case data := <-client.Ch():
				fmt.Printf("Receive: %s\n", data)
			}
		}
	})

	client.OnClose(func() {
		fmt.Printf("Close Server!!")
	})
``````
