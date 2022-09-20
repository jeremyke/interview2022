package main

import "fmt"

//给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。

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
	aLink := &ListNode{Val: 2}
	addNode(aLink, &ListNode{Val: 4})
	//addNode(aLink, &ListNode{Val: 5})
	bLink := &ListNode{Val: 1}
	addNode(bLink, &ListNode{Val: 9})
	addNode(bLink, &ListNode{Val: 1})
	addNode(bLink, aLink)
	cLink := &ListNode{Val: 3}
	//addNode(cLink, &ListNode{Val: 6})
	//addNode(cLink, &ListNode{Val: 1})
	addNode(cLink, aLink)
	showLink(bLink)
	showLink(cLink)
	res := getIntersectionNode(bLink, cLink)
	showLink(res)
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	comHeadMap := make(map[*ListNode]int, 2)
	for headA != nil {
		comHeadMap[headA]++
		headA = headA.Next
	}
	for headB != nil {
		if _, ok := comHeadMap[headB]; ok {
			return headB
		}
		comHeadMap[headA]++
		headB = headB.Next
	}
	return nil
}
