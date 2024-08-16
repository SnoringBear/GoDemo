package collection

import (
	"fmt"
	"testing"
)

func TestSlice01(t *testing.T) {
	// 初始化一个长度为10的数组,元素为零值
	a := make([]int, 10)
	for i := 0; i < 12; i++ {
		a = append(a, i)
		fmt.Printf("lenth = %d,cap = %d\n", len(a), cap(a))
	}
	fmt.Println(a)
}

func TestSlice02(t *testing.T) {
	// 声明切片
	var a []int
	a = append(a, 1)
	fmt.Println(a)
}
