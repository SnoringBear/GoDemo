package net_demo

import (
	"fmt"
	"net"
	"os"
	"testing"
)

func TestTCP01(t *testing.T) {
	// 设置目标地址和端口
	server := "localhost:9999"

	// 解析目标地址
	addr, err := net.ResolveTCPAddr("tcp", server)
	if err != nil {
		fmt.Println("Error resolving address:", err)
		os.Exit(1)
	}

	// 发起TCP连接
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		fmt.Println("Error dialing:", err)
		os.Exit(1)
	}

	defer func(conn *net.TCPConn) {
		err := conn.Close()
		if err != nil {
			os.Exit(1)
		}
	}(conn) // 连接结束时关闭

	// 向服务器发送数据
	message := "send message!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error writing to server:", err)
		os.Exit(1)
	}
	fmt.Println("Message sent:", message)

	// 接收服务器的响应
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading from server:", err)
		os.Exit(1)
	}
	fmt.Println("Received from server:", string(buf[:n]))
}
