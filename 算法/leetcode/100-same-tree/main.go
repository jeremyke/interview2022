package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	root2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 2}}
	//root := &TreeNode{Val: 0}
	fmt.Println(isSameTree(root1, root2))
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	pArr := serializeByPreSort(p)
	qArr := serializeByPreSort(q)
	if len(pArr) == 0 && len(qArr) == 0 {
		return true
	}
	if len(pArr) != len(qArr) {
		return false
	}
	for i := 0; i < len(pArr); i++ {
		if qArr[i] != pArr[i] {
			return false
		}
	}
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
