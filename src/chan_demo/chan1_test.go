package chan_demo

import (
	"fmt"
	"testing"
	"time"
)

func TestCh01(t *testing.T) {
	// 创建一个带有缓冲区的 channel，容量为 3
	ch := make(chan int, 3)

	// 启动一个新的 goroutine 作为发送方
	go func() {
		// 发送 3 个值到 channel
		ch <- 10
		ch <- 20
		ch <- 30
		fmt.Println("发送方：所有值已发送完毕，即将关闭 channel。")

		// 关闭 channel，向接收方发出信号
		close(ch)
	}()

	fmt.Println("接收方：准备开始接收值...")

	// 使用 for...range 循环从 channel 接收值
	// 即使发送方的 goroutine 已经关闭了 channel，
	// 这个循环依然会继续，直到取完所有缓冲的值。
	for value := range ch {
		fmt.Printf("接收方：接收到值 %d\n", value)
		// 等待一秒，以便更清晰地观察输出顺序
		time.Sleep(1 * time.Second)
	}

	fmt.Println("接收方：Channel 已关闭且为空，循环结束。")
}
