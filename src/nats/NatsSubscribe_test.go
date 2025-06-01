package nats

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"testing"
)

func TestNatsSub01(t *testing.T) {
	nc, err := nats.Connect(nats.DefaultURL) // 默认地址是 nats://127.0.0.1:4222
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// 订阅主题 "updates"
	_, err = nc.Subscribe("updates", func(m *nats.Msg) {
		fmt.Printf("Received message: %s\n", string(m.Data))
	})
	if err != nil {
		log.Fatal(err)
	}

	// 保持运行，监听消息
	select {}
}
