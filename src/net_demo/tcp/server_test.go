package tcp

import (
	"errors"
	log2 "github.com/rs/zerolog/log"
	"net"
	"os"
	"testing"
)

func TestTCPServer01(t *testing.T) {
	address := "localhost:8080"
	listen, err := net.Listen("tcp", address)
	if err != nil {
		log2.Fatal().Msgf("Error listening:%v", err)
	}

	// 类似Java中 try-with-resources,不过会更加优雅
	defer func(listen net.Listener) {
		err := listen.Close()
		if err != nil {
			log2.Fatal().Msgf("Error closing listener:%v", err)
		}
	}(listen)

	for {
		conn, err := listen.Accept()
		if err != nil {
			log2.Warn().Msgf("Error accepting:%v", err)
			continue
		}
		log2.Info().Msgf("Accepted connection from %s", conn.RemoteAddr())
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log2.Warn().Msgf("Error closing connection:%v", err)
		}
	}(conn)
	buffer := make([]byte, 1024)
	for {
		// 从连接中读取数据到缓冲区
		n, err := conn.Read(buffer)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				log2.Fatal().Msg("Error: Connection closed by peer.")
			} else {
				log2.Fatal().Msgf("Error reading from server:%v", err)
			}
			break // 读取错误时退出循环
		}

		// 处理读取到的数据
		// 注意：n 是实际读取到的字节数，可能小于缓冲区的长度
		data := buffer[:n] // 截取实际读取到的数据部分
		log2.Info().Msgf("Received from server: %s", string(data))
		//conn.Write(data)
	}
}
