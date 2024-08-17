package collection

import (
	"fmt"
	"testing"
)

func TestSubList01(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}
	b := a[1:]
	fmt.Println(b)
}
