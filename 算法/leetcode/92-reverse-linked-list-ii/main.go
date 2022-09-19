package main

import "fmt"

//给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。
//请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。

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
	//addNode(aLink, &ListNode{Val: 4})
	//addNode(aLink, &ListNode{Val: 5})
	showLink(aLink)
	res := reverseBetween(aLink, 3, 3)
	showLink(res)
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var (
		smallList, rangeList, bigList, totalList = &ListNode{}, &ListNode{}, &ListNode{}, &ListNode{}
		smallTail, rangeTail, bigTail, totalTail = smallList, rangeList, bigList, totalList
		nodeArr                                  = make([]*ListNode, 0, right-left+1)
		index                                    = 1
	)

	for head != nil {
		//小于区间的
		if index < left {
			smallTail.Next = &ListNode{Val: head.Val}
			smallTail = smallTail.Next
		}
		//大于区间的
		if index > right {
			bigTail.Next = &ListNode{Val: head.Val}
			bigTail = bigTail.Next
		}
		//介于中间的
		if index >= left && index <= right {
			nodeArr = append(nodeArr, &ListNode{Val: head.Val})
		}
		index++
		head = head.Next
	}
	//把区间节点数组转为链表
	for i := len(nodeArr) - 1; i >= 0; i-- {
		rangeTail.Next = nodeArr[i]
		rangeTail = rangeTail.Next
	}
	//把三个区间串起来
	smallList = smallList.Next
	rangeList = rangeList.Next
	bigList = bigList.Next
	if smallList != nil {
		totalTail.Next = smallList
		totalTail = smallTail
	}
	if rangeList != nil {
		totalTail.Next = rangeList
		totalTail = rangeTail
	}
	if bigList != nil {
		totalTail.Next = bigList
		totalTail = bigTail
	}
	return totalList.Next
}
