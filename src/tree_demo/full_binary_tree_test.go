package tree_demo

import (
	"fmt"
	"testing"
)

// 定义二叉树节点结构
type Node2 struct {
	Value       int
	Left, Right *Node2
}

// 创建一个新节点
func NewNode2(value int) *Node2 {
	return &Node2{Value: value}
}

// 插入节点（确保是满二叉树的结构）
func Insert(root *Node2, value int) *Node2 {
	if root == nil {
		return NewNode2(value)
	}

	// 如果左子树为空，插入到左子树
	if root.Left == nil {
		root.Left = NewNode2(value)
		return root
	}

	// 如果右子树为空，插入到右子树
	if root.Right == nil {
		root.Right = NewNode2(value)
		return root
	}

	// 如果左右子树都不为空，递归地插入到下一层
	if root.Left != nil && root.Right != nil {
		Insert(root.Left, value)
	}

	return root
}

// 打印二叉树（前序遍历）
func PreOrderTraversal(root *Node2) {
	if root != nil {
		fmt.Print(root.Value, " ")
		PreOrderTraversal(root.Left)
		PreOrderTraversal(root.Right)
	}
}

func TestNode201(t *testing.T) {
	root := NewNode2(1)
	Insert(root, 2)
	Insert(root, 3)
	Insert(root, 4)
	Insert(root, 5)

	fmt.Println("前序遍历：")
	PreOrderTraversal(root)
}
