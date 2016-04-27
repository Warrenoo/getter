# getter
- Sample WebSocket Client

### INSTALL
``````
  go get github.com/warrenoo/getter
``````

### USE
``````ruby
	var url string = "127.0.0.1:6000"
	var path string = "/websocket"

	client := getter.New(url, path)

	client.OnOpen(func() {
		fmt.Printf("Init: %s\n", url)
	})

	message := []byte("{'test':1}")
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
