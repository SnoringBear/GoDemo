package string_demo

import (
	"fmt"
	"testing"
	"time"
)

func TestFormat01(t *testing.T) {
	// 通过 fmt.Sprintf函数格式化出来的字符串,包含了"\n"
	timeStr := fmt.Sprintf("测试时间:%d \n", time.Now().Unix())
	// 只有[fmt.Print]、[fmt.Printf]函数处理了"\n"
	fmt.Print(timeStr)
}

func TestFormat02(t *testing.T) {
	fmt.Printf("测试时间:%d \n", time.Now().Unix())
}
