package collection

import (
	"fmt"
	"testing"
)

func TestFilter01(t *testing.T) {
	// 遍历并追加
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var filtered []int
	for _, num := range numbers {
		if num%2 == 0 {
			filtered = append(filtered, num)
		}
	}
	fmt.Println(filtered) // 输出：[2 4 6 8]
}

func TestFilter02(t *testing.T) {
	// 使用 filter 函数（自定义）
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	evenNumbers := filter(numbers, func(num int) bool {
		return num%2 == 0
	})
	fmt.Println(evenNumbers)
}

func filter(slice []int, f func(int) bool) []int {
	var result []int
	for _, v := range slice {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}
