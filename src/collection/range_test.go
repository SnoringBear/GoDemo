package collection

import (
	"fmt"
	"testing"
)

func Test01(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	for index, item := range a {
		fmt.Printf("index = %d, item = %d\n", index, item)
	}
}
