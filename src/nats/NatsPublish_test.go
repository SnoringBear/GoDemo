package nats

import (
	"github.com/nats-io/nats.go"
	"log"
	"testing"
)

func TestNatsPub01(t *testing.T) {
	// 连接到NATS服务器
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatalf("无法连接到NATS服务器: %v", err)
	}
	defer nc.Close()

	// 发布消息到主题
	subject := "example"
	message := []byte("Hello NATS!")

	err = nc.Publish(subject, message)
	if err != nil {
		log.Fatalf("无法发布消息: %v", err)
	}

	log.Println("消息已发送")
}
