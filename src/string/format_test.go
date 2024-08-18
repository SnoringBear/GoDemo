package string

import (
	"fmt"
	"testing"
	"time"
)

func TestFormat01(t *testing.T) {
	timeStr := fmt.Sprintf("测试时间:%d", time.Now().Unix())
	fmt.Println(timeStr)
}
