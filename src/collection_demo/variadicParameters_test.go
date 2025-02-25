package collection_demo

import (
	"github.com/rs/zerolog/log"
	"testing"
)

// 在 Go 语言中，... 语法主要用于处理可变参数（variadic parameters）和切片操作。以下是这两种用法的详细说明：

// 可变参数允许函数接受不定数量的参数。在函数定义中，通过在参数类型前加上 ... 来表示该参数是可变的。在函数体内，这个参数被视为一个切片。

// TestVariadicParam01   variadic param   多个参数
func TestVariadicParam01(t *testing.T) {
	i := sum(1, 2, 3, 4)
	log.Info().Msgf("Sum is: %d", i)
}

// TestVariadicParam02 切片变成可变参数
func TestVariadicParam02(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	i := sum(s...)
	log.Info().Msgf("Sum is: %d", i)
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
