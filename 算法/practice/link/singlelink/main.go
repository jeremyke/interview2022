// Package main
// @Description: 单向链表演示
package main

import "fmt"

//
// LinkNode
// @Description: 链表节点
//
type LinkNode struct {
	value string
	next  *LinkNode
}

//
// AddNode
// @Description: 添加节点
// @Author: maxwell.ke
// @time 2022-10-26 13:59:34
// @receiver ln
// @param n
// @return error
//
func (ln *LinkNode) AddNode(addNode *LinkNode) error {
	tail := ln
	for {
		if tail.next == nil {
			//找到了最后一个结点
			break
		}
		//如果没有找到，把当前结点的next指向临时结点
		tail = tail.next
	}
	tail.next = addNode
	return nil
}

//
// DelNode
// @Description: 移除节点
// @Author: maxwell.ke
// @time 2022-10-26 14:07:13
// @receiver ln
// @param delNode
// @return error
//
func (ln *LinkNode) DelNode(delNode *LinkNode) error {
	tail := ln
	for {
		//找到比需要删除的节点的前一个节点
		if delNode == tail { //删除的是头节点,不可删除
			return fmt.Errorf("头节点不可删除")
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
	return nil
}

//
// ShowLink
// @Description: 遍历链表
// @Author: maxwell.ke
// @time 2022-10-26 14:30:51
// @receiver ln
//
func (ln *LinkNode) ShowLink() {
	tail := ln
	for {
		if tail == nil {
			break
		}
		fmt.Printf("[%s]->", tail.value)
		tail = tail.next
	}
	fmt.Println()
}

func main() {
	head := &LinkNode{}
	node1 := &LinkNode{value: "aa"}
	node2 := &LinkNode{value: "bb"}
	node3 := &LinkNode{value: "cc"}
	node4 := &LinkNode{value: "dd"}
	node5 := &LinkNode{value: "ee"}
	head.AddNode(node1)
	head.AddNode(node2)
	head.AddNode(node3)
	head.AddNode(node4)
	head.AddNode(node5)
	head.ShowLink()
	head.DelNode(node3)
	head.ShowLink()
}
