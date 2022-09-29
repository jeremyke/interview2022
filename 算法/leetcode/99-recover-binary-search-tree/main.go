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
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 3, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 4}}
	//root := &TreeNode{Val: 0}
	recoverTree(root)
}

//思路：
//1.按照中序遍历转为数组
//2.数组排序
//3.数组转为二叉树

func recoverTree(root *TreeNode) {
	str, arr := treeTransformStrArr(root)
	fmt.Println(str)
	fmt.Println(arr)
	//sort.Ints(arr)

}

//中序遍历二叉树序列化，且把非空节点存入数组
func treeTransformStrArr(root *TreeNode) (str string, arr []int) {
	var process func(root *TreeNode)
	process = func(root *TreeNode) {
		if root == nil {
			str = fmt.Sprint(str, "#_")
			return
		}
		if root.Left != nil {
			process(root.Left)
		}
		arr = append(arr, root.Val)
		str = fmt.Sprint(str, root.Val, "_")
		if root.Right != nil {
			process(root.Right)
		}
	}
	process(root)
	return
}

//数组转为中序遍历二叉树
//func arrTransformTree(arr []int) *TreeNode {
//	var tree *TreeNode
//	for k, v := range arr {
//
//	}
//}
