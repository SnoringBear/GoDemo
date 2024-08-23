package gonet

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
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

func TestGrpc03(t *testing.T) {
	// 创建 gRPC 服务器
	grpcServer := grpc.NewServer()

	// 注册服务
	// pb.RegisterMyServiceServer(grpcServer, &myService{})

	// 启动监听器
	fmt.Println("1111111111111111111111111111111111")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("22222222222222222222222222222222222")
	// 启动 gRPC 服务器
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
