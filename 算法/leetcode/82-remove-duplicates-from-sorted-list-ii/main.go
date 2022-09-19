package main

import "fmt"

//给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。

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
	addNode(aLink, &ListNode{Val: 2})
	addNode(aLink, &ListNode{Val: 4})
	addNode(aLink, &ListNode{Val: 5})
	showLink(aLink)
	res := deleteDuplicates(aLink)
	showLink(res)
}

func deleteDuplicates(head *ListNode) *ListNode {
	var (
		copyHead     = head
		nodeValueMap = make(map[int]int, 1)
		newList      = &ListNode{}
		tail         = newList
	)
	for head != nil {
		nodeValueMap[head.Val]++
		head = head.Next
	}

	for copyHead != nil {
		if nodeValueMap[copyHead.Val] == 1 {
			tail.Next = &ListNode{Val: copyHead.Val}
			tail = tail.Next
		}
		copyHead = copyHead.Next
	}
	return newList.Next
}
