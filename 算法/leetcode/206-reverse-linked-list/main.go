package main

import "fmt"

//给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

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
	//addNode(aLink, &ListNode{Val: 1})
	//addNode(aLink, &ListNode{Val: 4})
	//addNode(aLink, &ListNode{Val: 5})
	showLink(aLink)
	res := reverseList(aLink)
	showLink(res)
}

func reverseList(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != nil {
		curTmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = curTmp
	}
	return pre
}
