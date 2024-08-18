package generics

import (
	"fmt"
	"testing"
)

func TestGenerics01(t *testing.T) {
	// 在 Go 1.18 版本开始，Go 语言引入了泛型（Generics）的支持
	// 泛型允许你编写能够处理任意类型的代码，而无需为每个具体类型编写单独的实现。
	// 这是通过引入类型参数的方式实现的，类似于其他静态类型语言中的泛型特性，如 C++、Java 和 Rust。
	Print("TestGenerics01")
}

// 泛型用于方法
func Print[T any](value T) {
	// any 是golang 1.18引入的语法糖,本质上是interface{}的别名
	// 类型参数使用方括号 [] 来定义
	fmt.Println(value)
}

func TestGenerics02(t *testing.T) {
	add := Add(2, 3)
	fmt.Println(add)
}

// 泛型约束
func Add[T int | float64](a, b T) T {
	return a + b
}

func TestGenerics03(t *testing.T) {
	p := Pair[int]{first: 1, second: 2}
	fmt.Println(p)
}

// 泛型用于结构体
type Pair[T any] struct {
	first, second T
}

func TestGenerics04(t *testing.T) {
	var intStack Stack[int]
	intStack.Push(10)
	intStack.Push(20)
	fmt.Println(intStack.Pop()) // 输出: 20 true

	var stringStack Stack[string]
	stringStack.Push("hello")
	stringStack.Push("world")
	fmt.Println(stringStack.Pop()) // 输出: world true
}

// 定义泛型栈
type Stack[T any] struct {
	elements []T
}

// Push 方法
func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

// Pop 方法
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		var zeroValue T
		return zeroValue, false
	}
	index := len(s.elements) - 1
	element := s.elements[index]
	s.elements = s.elements[:index]
	return element, true
}
