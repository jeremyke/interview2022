package main

import "fmt"

//给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。

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
	res := rotateRight(aLink, 2)
	showLink(res)
}

func rotateRight(head *ListNode, k int) *ListNode {
	//链表存入数组
	var (
		listCount int
		copyHead  = head
		newList   = &ListNode{}
		tail      = newList
	)
	for copyHead != nil {
		listCount++
		copyHead = copyHead.Next
	}
	nodeArr := make([]*ListNode, listCount)
	i := 0
	for head != nil {
		index := (i + k) % listCount
		nodeArr[index] = &ListNode{Val: head.Val}
		head = head.Next
		i++
	}
	//把数组元素取出来串成链表
	for _, value := range nodeArr {
		tail.Next = value
		tail = tail.Next
	}
	return newList.Next
}
