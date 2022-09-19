package main

import "fmt"

//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

func main() {
	head := &ListNode{
		Val: 1,
	}
	addNode(head, &ListNode{Val: 2})
	addNode(head, &ListNode{Val: 3})
	addNode(head, &ListNode{Val: 4})
	addNode(head, &ListNode{Val: 5})
	showLink(head)
	newLink := removeNthFromEnd(head, 4)
	showLink(newLink)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var (
		nodeArr = make([]*ListNode, 0, 1)
		tmpList = head
		newHead *ListNode
		tail    *ListNode
	)
	//原链表转为数组
	for tmpList != nil {
		nodeArr = append(nodeArr, tmpList)
		tmpList = tmpList.Next
	}

	//便历数组，去掉需要去掉的节点，连接起来成一个新的链表
	for k, v := range nodeArr {
		if k == len(nodeArr)-n {
			continue
		}
		v.Next = nil
		if newHead == nil {
			newHead = v
			tail = newHead
		} else {
			tail.Next = v
			tail = tail.Next
		}

	}
	return newHead
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

//func delNode(head, delNode *ListNode) error {
//	tempNode := head
//	for {
//		//找到比需要删除的节点的前一个节点
//		if delNode.no == 0 {
//			return fmt.Errorf("头结点不可删除")
//		}
//		if tempNode.next == nil {
//			return fmt.Errorf("没有找到需要删除的节点")
//		}
//		if tempNode.next.no == delNode.no {
//			break
//		}
//		//如果没有找到，把当前结点的next指向临时结点
//		tempNode = tempNode.next
//	}
//	//开始删除
//	tempNode.next = delNode.next
//	return nil
//}

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
