package main

import "fmt"

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func middleSearch(root *TreeNode, pre, post func(*TreeNode)) {
	// left -> root -> right
	if root == nil {
		return
	}

	middleSearch(root.Left, pre, post)
	pre(root)
	fmt.Printf("now is middle %v \n", root.Value)
	post(root)
	middleSearch(root.Right, pre, post)
}

func pre(root *TreeNode) {
	fmt.Printf("pre %v \n", root.Value)
}

func post(root *TreeNode) {
	fmt.Printf("post %v \n", root.Value)
}

func main() {
	// 这里可以创建一个二叉树并进行操作
	var root *TreeNode
	root = &TreeNode{Value: 1}
	root.Left = &TreeNode{Value: 2}
	root.Right = &TreeNode{Value: 3}
	middleSearch(root, pre, post)
}
