package gonet

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/url"
	"testing"
)

func TestWeb01(t *testing.T) {
	http.HandleFunc("/ws", wsHandler)

	fmt.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许所有请求来源，生产环境中可以进行限制
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// 将 HTTP 连接升级为 WebSocket 连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading:", err)
		return
	}
	defer conn.Close()

	for {
		// 读取客户端消息
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}

		// 打印消息
		fmt.Printf("Received: %s\n", message)

		// 将消息发送回客户端
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			fmt.Println("Error writing message:", err)
			break
		}
	}
}

func TestWeb02(t *testing.T) {
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws"}
	fmt.Println("Connecting to", u.String())

	// 创建 WebSocket 连接
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	defer conn.Close()

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
