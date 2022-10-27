package main

import "fmt"

type LinkNode struct {
	value string
	pre   *LinkNode
	next  *LinkNode
}

//
// AddNode
// @Description: 添加节点
// @Author: maxwell.ke
// @time 2022-10-26 14:47:52
// @receiver ln
// @param node
//
func (ln *LinkNode) AddNode(node *LinkNode) {
	tail := ln
	for {
		if tail.next == nil {
			//找到了最后一个结点
			break
		}
		//如果没有找到，把当前结点的next指向临时结点
		tail = tail.next
	}
	tail.next = node
	node.pre = tail
}

//
// DelNode
// @Description: 删除节点
// @Author: maxwell.ke
// @time 2022-10-26 14:51:27
// @receiver ln
// @param delNode
// @return error
//
func (ln *LinkNode) DelNode(delNode *LinkNode) error {
	tail := ln
	for {
		//找到比需要删除的节点的前一个节点
		if delNode == ln {
			return fmt.Errorf("头结点不可删除")
		}
		if tail.next == nil {
			return fmt.Errorf("没有找到需要删除的节点")
		}
		if tail.next == delNode {
			break
		}
		//如果没有找到，把当前结点的next指向临时结点
		tail = tail.next
	}
	//开始删除
	tail.next = delNode.next
	delNode.next.pre = tail
	return nil
}

//
// ShowLink
// @Description: 遍历链表
// @Author: maxwell.ke
// @time 2022-10-26 14:54:17
// @receiver ln
//
func (ln *LinkNode) ShowLink() {
	tail := ln
	if tail.next == nil {
		fmt.Println("空链表")
		return
	}
	for {
		fmt.Printf("[%s]->", tail.next.value)
		tail = tail.next
		if tail.next == nil {
			break
		}
	}
	fmt.Println()
}
