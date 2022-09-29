package main

import "fmt"

//给你一个整数n ，求恰由n个节点组成且节点值从1到n互不相同的二叉搜索树有多少种？
//返回满足题意的二叉搜索树的种数。

func main() {
	fmt.Println(numTrees(3))
}

func numTrees(n int) int {
	return process(1, n)
}

var count int

func process(start, end int) int {
	if start > end {
		return 1
	}
	//枚举可行的根节点
	for i := start; i <= end; i++ {
		//左子树集合
		leftTrees := process(start, i-1)
		//右子树集合
		rightTrees := process(i+1, end)
		//遍历左右子树集合
		count += leftTrees * rightTrees

	}
	return count
}
