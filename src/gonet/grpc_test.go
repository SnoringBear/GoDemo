package gonet

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"testing"
)

func TestGrpc01(t *testing.T) {
	// 创建 gRPC 服务器
	grpcServer := grpc.NewServer()
	fmt.Println(grpcServer)
}

func TestGrpc02(t *testing.T) {
	// 加载服务端证书和密钥
	creeds, err := credentials.NewServerTLSFromFile("server.crt", "server.key")
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}

	// 启用 TLS 的 gRPC 服务器
	grpcServer := grpc.NewServer(grpc.Creds(creeds))
	fmt.Println(grpcServer)
}
