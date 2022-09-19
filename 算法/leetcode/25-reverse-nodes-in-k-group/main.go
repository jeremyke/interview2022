package main

import (
	"fmt"
)

//给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
//
//k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//
//你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
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
	addNode(aLink, &ListNode{Val: 4})
	addNode(aLink, &ListNode{Val: 5})
	showLink(aLink)
	res := reverseKGroup(aLink, 3)
	showLink(res)
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	newList := &ListNode{}
	tail := newList
	for head != nil {
		//节点放入数组
		listDirection := 0
		nodeArr := make([]*ListNode, 0, k)
		for i := 1; i <= k; i++ {
			if head == nil {
				listDirection = 1
				break
			}
			nodeArr = append(nodeArr, &ListNode{Val: head.Val})
			head = head.Next
		}
		//节点从数组中取出
		if listDirection == 1 { //节点从数组正向取出
			for _, value := range nodeArr {
				tail.Next = value
				tail = tail.Next
			}
		} else if listDirection == 0 { //节点从数组向逆取出
			for i := len(nodeArr) - 1; i >= 0; i-- {
				tail.Next = nodeArr[i]
				tail = tail.Next
			}
		}

	}
	return newList.Next
}
