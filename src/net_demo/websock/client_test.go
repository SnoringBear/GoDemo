package websock

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"testing"
)

func TestWeb02(t *testing.T) {
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws"}
	fmt.Println("Connecting to", u.String())

	// 创建 WebSocket 连接
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatal("Close error:", err)
		}
	}(conn)

	// 发送消息给服务器
	err = conn.WriteMessage(websocket.TextMessage, []byte("Hello, Server!"))
	if err != nil {
		log.Println("Write error:", err)
		return
	}

	// 读取服务器响应
	_, message, err := conn.ReadMessage()
	if err != nil {
		log.Println("Read error:", err)
		return
	}
	fmt.Printf("Received from server: %s\n", message)
}
