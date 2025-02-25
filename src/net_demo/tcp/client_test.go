package tcp

import (
	log2 "github.com/rs/zerolog/log"
	"net"
	"testing"
)

func TestTCPClient01(t *testing.T) {
	// 设置目标地址和端口
	server := "localhost:8081"

	// 解析目标地址
	addr, err := net.ResolveTCPAddr("tcp", server)
	if err != nil {
		log2.Fatal().Msgf("Error resolving address:%v", err)
	}

	// 发起TCP连接
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log2.Fatal().Msgf("Error connecting to server:%v", err)
	}

	defer func(conn *net.TCPConn) {
		err := conn.Close()
		if err != nil {
			log2.Fatal().Msgf("Error closing connection:%v", err)
		}
	}(conn) // 连接结束时关闭

	// 向服务器发送数据
	message := "send message!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		log2.Fatal().Msgf("Error writing to server:%v", err)
	}
	log2.Warn().Msgf("Message sent:%s", message)

	for {
		// 接收服务器的响应
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			log2.Fatal().Msgf("Error reading from server:%v", err)
		}
		log2.Info().Msgf("Received from server:%s", string(buf[:n]))

		_, err = conn.Write([]byte(message))
		if err != nil {
			log2.Fatal().Msgf("Error writing to server:%v", err)
		}
	}

}
