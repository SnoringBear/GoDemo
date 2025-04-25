package b_plus_tree

import (
	"encoding/gob"
	"os"
	"sort"
)

const M = 4

type Node struct {
	IsLeaf   bool
	Keys     []int
	Children []*Node
	Values   []string
	Next     *Node
}

type BPlusTree struct {
	Root *Node
}

func NewBPlusTree() *BPlusTree {
	return &BPlusTree{
		Root: &Node{IsLeaf: true},
	}
}

func (t *BPlusTree) Search(key int) (string, bool) {
	cur := t.Root
	for !cur.IsLeaf {
		i := sort.Search(len(cur.Keys), func(i int) bool {
			return key < cur.Keys[i]
		})
		cur = cur.Children[i]
	}
	for i, k := range cur.Keys {
		if k == key {
			return cur.Values[i], true
		}
	}
	return "", false
}

func (t *BPlusTree) Insert(key int, value string) {
	root := t.Root
	newKey, newChild := t.insertRecursive(root, key, value)
	if newChild != nil {
		t.Root = &Node{
			IsLeaf:   false,
			Keys:     []int{newKey},
			Children: []*Node{root, newChild},
		}
	}
}

func (t *BPlusTree) insertRecursive(node *Node, key int, value string) (int, *Node) {
	if node.IsLeaf {
		i := sort.SearchInts(node.Keys, key)
		node.Keys = append(node.Keys[:i], append([]int{key}, node.Keys[i:]...)...)
		node.Values = append(node.Values[:i], append([]string{value}, node.Values[i:]...)...)
		if len(node.Keys) < M {
			return 0, nil
		}
		return t.splitLeaf(node)
	}
	i := sort.Search(len(node.Keys), func(i int) bool {
		return key < node.Keys[i]
	})
	newKey, newChild := t.insertRecursive(node.Children[i], key, value)
	if newChild != nil {
		node.Keys = append(node.Keys[:i], append([]int{newKey}, node.Keys[i:]...)...)
		node.Children = append(node.Children[:i+1], append([]*Node{newChild}, node.Children[i+1:]...)...)
		if len(node.Keys) < M {
			return 0, nil
		}
		return t.splitInternal(node)
	}
	return 0, nil
}

func (t *BPlusTree) splitLeaf(node *Node) (int, *Node) {
	mid := M / 2
	newNode := &Node{
		IsLeaf: true,
		Keys:   append([]int(nil), node.Keys[mid:]...),
		Values: append([]string(nil), node.Values[mid:]...),
		Next:   node.Next,
	}
	node.Keys = node.Keys[:mid]
	node.Values = node.Values[:mid]
	node.Next = newNode
	return newNode.Keys[0], newNode
}

func (t *BPlusTree) splitInternal(node *Node) (int, *Node) {
	mid := M / 2
	newNode := &Node{
		IsLeaf:   false,
		Keys:     append([]int(nil), node.Keys[mid+1:]...),
		Children: append([]*Node(nil), node.Children[mid+1:]...),
	}
	upKey := node.Keys[mid]
	node.Keys = node.Keys[:mid]
	node.Children = node.Children[:mid+1]
	return upKey, newNode
}

func (t *BPlusTree) Delete(key int) {
	t.deleteRecursive(t.Root, key)
	if !t.Root.IsLeaf && len(t.Root.Keys) == 0 {
		t.Root = t.Root.Children[0]
	}
}

func (t *BPlusTree) deleteRecursive(node *Node, key int) bool {
	if node.IsLeaf {
		for i, k := range node.Keys {
			if k == key {
				node.Keys = append(node.Keys[:i], node.Keys[i+1:]...)
				node.Values = append(node.Values[:i], node.Values[i+1:]...)
				return true
			}
		}
		return false
	}
	i := sort.Search(len(node.Keys), func(i int) bool {
		return key < node.Keys[i]
	})
	deleted := t.deleteRecursive(node.Children[i], key)
	if deleted && len(node.Children[i].Keys) == 0 {
		node.Children = append(node.Children[:i], node.Children[i+1:]...)
		if i > 0 {
			node.Keys = append(node.Keys[:i-1], node.Keys[i:]...)
		} else {
			node.Keys = node.Keys[1:]
		}
	}
	return deleted
}

func (t *BPlusTree) RangeSearch(start, end int) []string {
	result := []string{}
	cur := t.Root
	for !cur.IsLeaf {
		i := sort.Search(len(cur.Keys), func(i int) bool {
			return start < cur.Keys[i]
		})
		cur = cur.Children[i]
	}
	for cur != nil {
		for i, k := range cur.Keys {
			if k >= start && k <= end {
				result = append(result, cur.Values[i])
			} else if k > end {
				return result
			}
		}
		cur = cur.Next
	}
	return result
}

func (t *BPlusTree) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	enc := gob.NewEncoder(file)
	return enc.Encode(t)
}

func LoadFromFile(filename string) (*BPlusTree, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	dec := gob.NewDecoder(file)
	var tree BPlusTree
	if err := dec.Decode(&tree); err != nil {
		return nil, err
	}
	return &tree, nil
}

func init() {
	gob.Register(&BPlusTree{})
	gob.Register(&Node{})
}
