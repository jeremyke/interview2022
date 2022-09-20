package main

import "fmt"

//给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
//1-3-1
//2-2-3-3

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
	res := isPalindrome(aLink)
	fmt.Println(res)
}

func isPalindrome(head *ListNode) bool {
	var (
		leftList, rightList = &ListNode{}, &ListNode{}
		leftTail            = leftList
		fastTail, slowTail  = head, head
	)
	//左半边节点存入leftList
	for slowTail != nil && fastTail != nil {
		leftTail.Next = &ListNode{Val: slowTail.Val}
		leftTail = leftTail.Next

		//奇数链表fast.Next为nil,偶数链表的fast.Next,Next为nil
		slowTail = slowTail.Next
		if fastTail.Next != nil {
			fastTail = fastTail.Next.Next
		} else {
			break
		}
	}
	//反转右边节点
	rightList = reverseList(slowTail)

	//比较2个链表
	flag := true
	leftList = leftList.Next
	for leftList != nil && rightList != nil {
		if leftList.Val != rightList.Val {
			flag = false
			break
		}
		leftList = leftList.Next
		rightList = rightList.Next
	}
	return flag
}

func reverseList(head *ListNode) *ListNode {
	var prev, cur *ListNode = nil, head
	for cur != nil {
		nextTmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextTmp
	}
	return prev
}
