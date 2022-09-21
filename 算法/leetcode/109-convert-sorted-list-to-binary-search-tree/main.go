package main

import "fmt"

//给定一个单链表的头节点  head ，其中的元素 按升序排序 ，将其转换为高度平衡的二叉搜索树。
//
//本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差不超过 1。
//

type ListNode struct {
	Val  int
	Next *ListNode
}

func addNode(head, newNode *ListNode) {
	temp := head
	for {
		if temp.Next == nil {
			//找到了最后一个结点
			break
		}
		//如果没有找到，把当前结点的next指向临时结点
		temp = temp.Next
	}
	temp.Next = newNode
}

func showLink(head *ListNode) {
	temp := head
	for {
		if temp == nil {
			break
		}
		fmt.Printf("[%v]->", temp.Val)
		temp = temp.Next
	}
	fmt.Println()
}

func main() {
	aLink := &ListNode{Val: 1}
	addNode(aLink, &ListNode{Val: 2})
	addNode(aLink, &ListNode{Val: 3})
	//addNode(aLink, &ListNode{Val: 4})
	//addNode(aLink, &ListNode{Val: 5})
	showLink(aLink)
	res := reverseBetween(aLink, 3, 3)
	showLink(res)
}

func reverseBetween() {

}
