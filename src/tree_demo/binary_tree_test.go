package tree_demo

import (
	"fmt"
	"testing"
)

// Node 定义二叉树的节点结构
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Insert 插入新节点
func (n *Node) Insert(value int) {
	if value < n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: value}
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Value: value}
		} else {
			n.Right.Insert(value)
		}
	}
}

// InOrder 中序遍历
func (n *Node) InOrder() {
	if n == nil {
		return
	}
	n.Left.InOrder()
	fmt.Print(n.Value, " ")
	n.Right.InOrder()
}

// Min 查找最小值
func (n *Node) Min() *Node {
	if n.Left != nil {
		return n.Left.Min()
	}
	return n
}

// Max 查找最大值
func (n *Node) Max() *Node {
	if n.Right != nil {
		return n.Right.Max()
	}
	return n
}

// Delete 删除节点
func (n *Node) Delete(value int, parent *Node) *Node {
	if value < n.Value {
		if n.Left != nil {
			n.Left = n.Left.Delete(value, n)
		}
	} else if value > n.Value {
		if n.Right != nil {
			n.Right = n.Right.Delete(value, n)
		}
	} else {
		// 找到要删除的节点
		if n.Left == nil && n.Right == nil {
			return nil // 没有子节点
		} else if n.Left == nil {
			return n.Right // 只有右子节点
		} else if n.Right == nil {
			return n.Left // 只有左子节点
		} else {
			// 有两个子节点，找到右子树的最小值
			minNode := n.Right.Min()
			n.Value = minNode.Value
			n.Right = n.Right.Delete(minNode.Value, n)
		}
	}
	return n
}

func TestBinaryTree01(t *testing.T) {
	root := &Node{Value: 10}
	root.Insert(5)
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)
	root.Insert(13)
	root.Insert(17)

	fmt.Println("In-order traversal:")
	root.InOrder() // 中序遍历
	fmt.Println()

	fmt.Println("Deleting 7...")
	root.Delete(7, nil)
	root.InOrder()
}
