package main

import (
	"fmt"
	"math"
)

//给你一个二叉树的根节点root ，判断其是否是一个有效的二叉搜索树。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	//root := &TreeNode{Val: 0}
	fmt.Println(isValidBST(root))
}

func isValidBST(root *TreeNode) bool {
	var process func(root *TreeNode) bool
	var preValue int = math.MinInt64
	process = func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		leftBST := process(root.Left)
		if !leftBST {
			return leftBST
		}
		if root.Val <= preValue {
			return false
		} else {
			preValue = root.Val
		}
		rightBST := process(root.Right)
		if !rightBST {
			return rightBST
		}
		return true
	}
	return process(root)
}
