package time

import (
	"fmt"
	"testing"
	"time"
)

func TestTick01(t *testing.T) {
	// 定时器
	tick := time.Tick(time.Second)
	for i := range tick {
		fmt.Println(i)
	}
}
