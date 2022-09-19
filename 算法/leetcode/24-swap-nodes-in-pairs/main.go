package main

import "fmt"

//给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。

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
	addNode(aLink, &ListNode{Val: 4})
	addNode(aLink, &ListNode{Val: 5})
	showLink(aLink)
	res := swapPairs(aLink)
	showLink(res)
}

func swapPairs(head *ListNode) *ListNode {
	var (
		newList = &ListNode{}
		tail    = newList
	)
	for head != nil && head.Next != nil {
		leftNode := &ListNode{Val: head.Val}
		rightNode := &ListNode{Val: head.Next.Val}
		rightNode.Next = leftNode
		tail.Next = rightNode
		tail = tail.Next.Next
		head = head.Next.Next
	}
	if head != nil {
		tail.Next = head
	}
	return newList.Next
}
