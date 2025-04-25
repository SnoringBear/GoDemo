package range_demo

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	a := [3]int{1, 2, 3}
	for index, item := range a {
		fmt.Println(index, item)
	}
}

func TestSlice(t *testing.T) {
	nums := []int{2, 4, 6}
	for i, num := range nums {
		fmt.Println(i, num)
	}
}

func TestString(t *testing.T) {
	s := "hello world"
	for i, c := range s {
		fmt.Printf("%d: %c\n", i, c)
	}
}

func TestMap(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	for key, val := range m {
		fmt.Println(key, val)
	}
}

func TestChan(t *testing.T) {
	ch := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
		close(ch)
	}()

	// 会阻塞直到有数据，直到 ch 被关闭
	for v := range ch {
		fmt.Println(v)
	}
}
