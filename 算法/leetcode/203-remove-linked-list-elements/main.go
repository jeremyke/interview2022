package main

import "fmt"

//给你一个链表的头节点head和一个整数val ，请你删除链表中所有满足 Node.val == val 的节点，并返回新的头节点 。

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
	addNode(aLink, &ListNode{Val: 1})
	addNode(aLink, &ListNode{Val: 4})
	addNode(aLink, &ListNode{Val: 5})
	showLink(aLink)
	res := removeElements(aLink, 1)
	showLink(res)
}
func removeElements(head *ListNode, val int) *ListNode {
	var (
		newList = &ListNode{}
		tail    = newList
	)
	for head != nil {
		if head.Val != val {
			tail.Next = &ListNode{Val: head.Val}
			tail = tail.Next
		}
		head = head.Next

	}
	return newList.Next
}
