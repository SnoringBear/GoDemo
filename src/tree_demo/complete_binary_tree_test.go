package tree_demo

import "fmt"

// 定义二叉树的节点结构
type TreeNode struct {
	value       int
	left, right *TreeNode
}

// 完全二叉树结构体
type CompleteBinaryTree struct {
	root *TreeNode
}

// 插入节点函数（按层序遍历顺序插入）
func (tree *CompleteBinaryTree) Insert(value int) {
	newNode := &TreeNode{value: value}
	if tree.root == nil {
		tree.root = newNode
	} else {
		queue := []*TreeNode{tree.root}

		// 使用队列找到空缺位置插入新节点
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]

			if node.left == nil {
				node.left = newNode
				return
			} else {
				queue = append(queue, node.left)
			}

			if node.right == nil {
				node.right = newNode
				return
			} else {
				queue = append(queue, node.right)
			}
		}
	}
}

// 前序遍历（打印二叉树的节点值）
func (tree *CompleteBinaryTree) PreOrderTraversal(node *TreeNode) {
	if node != nil {
		fmt.Printf("%d ", node.value)
		tree.PreOrderTraversal(node.left)
		tree.PreOrderTraversal(node.right)
	}
}
