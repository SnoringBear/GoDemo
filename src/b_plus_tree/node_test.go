package b_plus_tree

import (
	"fmt"
	"testing"
)

func TestNode01(t *testing.T) {
	tree := NewBPlusTree()
	tree.Insert(5, "Five")
	tree.Insert(10, "Ten")
	tree.Insert(15, "Fifteen")
	tree.Insert(20, "Twenty")
	tree.Insert(25, "Twenty-five")

	tree.Delete(10)

	fmt.Println("Range 12~25:", tree.RangeSearch(12, 25))

	tree.SaveToFile("tree.gob")
	loaded, _ := LoadFromFile("tree.gob")
	if val, ok := loaded.Search(15); ok {
		fmt.Println("Loaded Tree Find 15:", val)
	}
}
