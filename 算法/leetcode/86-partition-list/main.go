package main

import "fmt"

//给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
//
//你应当 保留 两个分区中每个节点的初始相对位置。
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
	addNode(aLink, &ListNode{Val: 4})
	addNode(aLink, &ListNode{Val: 2})
	addNode(aLink, &ListNode{Val: 4})
	addNode(aLink, &ListNode{Val: 5})
	showLink(aLink)
	res := partition(aLink, 3)
	showLink(res)
}
func partition(head *ListNode, x int) *ListNode {
	var (
		smallList, bigList = &ListNode{}, &ListNode{}
		smallTail          = smallList
		bigTail            = bigList
	)
	for head != nil {
		if head.Val < x {
			smallTail.Next = &ListNode{Val: head.Val}
			smallTail = smallTail.Next
		} else {
			bigTail.Next = &ListNode{Val: head.Val}
			bigTail = bigTail.Next
		}
		head = head.Next
	}
	smallTail.Next = bigList.Next
	return smallList.Next
}
