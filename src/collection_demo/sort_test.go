package collection_demo

import (
	"fmt"
	"sort"
	"testing"
)

func TestSort01(t *testing.T) {
	// 内置排序
	var a []int
	a = append(a, 7, 5, 9, 1, 4, 2)
	sort.Ints(a)
	fmt.Println(a)
}

type SSlice []S

func TestSort02(t *testing.T) {
	// 自定义排序 需要实现sort.Interface接口
	list := make([]S, 0)
	list = append(list, S{a: 7}, S{a: 5}, S{a: 1}, S{a: 4}, S{a: 2})
	sort.Sort(SSlice(list))
	fmt.Println(list)
}

type S struct {
	a int
}

func (s SSlice) Len() int {
	return len(s)
}

func (s SSlice) Less(i, j int) bool {
	return s[i].a < s[j].a
}

func (s SSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
