package grpc

import (
	pb "GoDemo/src/proto_gen"
	"context"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
	"testing"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(_ context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func TestGrpcServer01(t *testing.T) {
	// 可选参数 可以设置为TLS加密的服务器
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}
	log.Info().Msgf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatal().Msgf("failed to serve: %v", err)
	}
}
