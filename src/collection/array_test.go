package collection

import "testing"

func TestArray01(t *testing.T) {
	a := [3]int{1, 2, 3}
	b := a
	b[0] = 100
	t.Log(a, b)
}
