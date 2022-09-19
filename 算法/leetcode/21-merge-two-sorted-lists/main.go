package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

func main() {
	aLink := &ListNode{Val: 1}
	addNode(aLink, &ListNode{Val: 2})
	addNode(aLink, &ListNode{Val: 4})
	showLink(aLink)
	bLink := &ListNode{Val: 1}
	addNode(bLink, &ListNode{Val: 3})
	addNode(bLink, &ListNode{Val: 4})
	showLink(bLink)
	cLink := mergeTwoLists(aLink, bLink)
	showLink(cLink)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	aNode, bNode := list1, list2
	var cLink *ListNode
	var tail *ListNode
	for aNode != nil && bNode != nil {
		//确定追加到新链表的节点
		var tmpNode *ListNode
		if aNode.Val <= bNode.Val {
			tmpNode = &ListNode{Val: aNode.Val}
			aNode = aNode.Next
		} else {
			tmpNode = &ListNode{Val: bNode.Val}
			bNode = bNode.Next
		}
		//加到新链表上
		if cLink == nil {
			cLink = tmpNode
			tail = tmpNode
		} else {
			tail.Next = tmpNode
			tail = tail.Next
		}
	}
	for aNode != nil {
		tmpNode := &ListNode{Val: aNode.Val}
		//加到新链表上
		if cLink == nil {
			cLink = tmpNode
			tail = tmpNode
		} else {
			tail.Next = tmpNode
			tail = tail.Next

		}
		aNode = aNode.Next
	}
	for bNode != nil {
		tmpNode := &ListNode{Val: bNode.Val}
		//加到新链表上
		if cLink == nil {
			cLink = tmpNode
			tail = tmpNode
		} else {
			tail.Next = tmpNode
			tail = tail.Next

		}
		bNode = bNode.Next
	}
	return cLink
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
