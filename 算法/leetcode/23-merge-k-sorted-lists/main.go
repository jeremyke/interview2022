package main

import "fmt"

//给你一个链表数组，每个链表都已经按升序排列。
//
//请你将所有链表合并到一个升序链表中，返回合并后的链表。

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
	//addNode(aLink, &ListNode{Val: 4})
	//addNode(aLink, &ListNode{Val: 5})
	//bLink := &ListNode{}
	//addNode(bLink, &ListNode{Val: 3})
	//addNode(bLink, &ListNode{Val: 4})
	cLink := &ListNode{Val: -1}
	//addNode(cLink, &ListNode{Val: 6})
	mergeList := mergeKLists([]*ListNode{aLink, nil, cLink})
	showLink(mergeList)
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

//循环合并两个链表的做法
//func mergeKLists(lists []*ListNode) *ListNode {
//	var resList *ListNode
//	for i := 0; i < len(lists); i++ {
//		resList = mergeTwoLists(resList, lists[i])
//	}
//	return resList
//}

//分治合并的做法
func mergeKLists(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	if l > r {
		return nil
	}
	mid := (l + r) / 2
	return mergeTwoLists(merge(lists, l, mid), merge(lists, mid+1, r))
}
