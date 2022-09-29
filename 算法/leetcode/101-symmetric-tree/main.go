package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}}
	//root2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 2}}
	//root := &TreeNode{Val: 0}
	fmt.Println(isSymmetric(root))
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var leftTree, rightTree string
	if root.Left != nil {
		leftTree = serializeByPreSort(root.Left)
	}
	if root.Right != nil {
		rightTree = serializeByPreSort(root.Right)
	}
	if len(leftTree) != len(rightTree) {
		return false
	}
	fmt.Println(leftTree, rightTree)
	return true

}

func serializeByPreSort(head *TreeNode) string {
	if head == nil {
		return "#_"
	}
	str := fmt.Sprint(head.Val, "_")
	str = fmt.Sprint(str, serializeByPreSort(head.Left))
	str = fmt.Sprint(str, serializeByPreSort(head.Right))
	return str
}
