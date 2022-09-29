package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}
	arr := inorderTraversal(root)
	fmt.Println(arr)

}

func inorderTraversal(root *TreeNode) (arr []int) {
	var process func(root *TreeNode)
	process = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			process(root.Left)
		}
		arr = append(arr, root.Val)
		if root.Right != nil {
			process(root.Right)
		}
	}
	process(root)
	return
}
