package main

import "fmt"

//给你一个整数n ，请你生成并返回所有由n个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树。
//可以按任意顺序返回答案。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	treeArr := generateTrees(3)
	for _, v := range treeArr {
		fmt.Println(serializeByPreSort(v))
	}
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

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return process(1, n)
}

func process(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	allTrees := []*TreeNode{}
	// 枚举可行根节点
	for i := start; i <= end; i++ {
		// 获得所有可行的左子树集合
		leftTrees := process(start, i-1)
		// 获得所有可行的右子树集合
		rightTrees := process(i+1, end)
		// 从左子树集合中选出一棵左子树，从右子树集合中选出一棵右子树，拼接到根节点上
		for _, left := range leftTrees {
			for _, right := range rightTrees {
				currTree := &TreeNode{i, nil, nil}
				currTree.Left = left
				currTree.Right = right
				allTrees = append(allTrees, currTree)
			}
		}
	}
	return allTrees
}
