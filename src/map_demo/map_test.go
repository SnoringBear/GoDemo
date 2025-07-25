package map_demo

import (
	"fmt"
	"testing"
)

func TestMap01(t *testing.T) {
	m := make(map[int]string)
	fmt.Printf("数值为:%s \n", m[1])
	// map_demo api真是少得可怜,相比Java map丰富操作的api
	// 判断是否存在这个key的方法
	s, ok := m[2]
	if ok {
		fmt.Printf("存在值 = %d的key,value = %s", 2, s)
	}
	// ok这种使用方法,还可以用于通道接受数值   s,ok<-ch
}

func TestMap02(t *testing.T) {
	m := map[int]int{
		1: 1,
		2: 2,
	}
	m[4] = 4
	delete(m, 3)
}
