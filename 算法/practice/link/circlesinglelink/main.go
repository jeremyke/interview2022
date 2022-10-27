package main

import "fmt"

type LinkNode struct {
	value string
	next  *LinkNode
}

//
// AddNode
// @Description: 添加链表
// @Author: maxwell.ke
// @time 2022-10-26 15:15:51
// @receiver ln
// @param newNode
//
func (ln *LinkNode) AddNode(newNode *LinkNode) {
	//头节点,直接加
	if ln.next == nil {
		ln.value = newNode.value
		ln.next = ln
		return
	}
	tail := ln
	for {
		if tail.next == ln {
			break
		}
		tail = tail.next
	}
	//加入链表
	tail.next = newNode
	newNode.next = ln
}

//
// DelNode
// @Description: 删除节点
// @Author: maxwell.ke
// @time 2022-10-26 15:51:54
// @receiver ln
// @param delNode
// @return newLink
// @return err
//
func (ln *LinkNode) DelNode(delNode *LinkNode) (newLink *LinkNode, err error) {
	//空链表
	if ln.next == nil {
		return nil, fmt.Errorf("链表为空")
	}

	//1个元素的链表
	if ln.next == ln {
		if ln == delNode {
			return nil, nil
		} else {
			return nil, fmt.Errorf("未找到要删除的节点")
		}
	}

	//多个元素的链表
	tailNode := ln //头节点
	rearNode := ln //尾节点
	//找到尾结点
	for {
		if rearNode.next == ln {
			break
		}
		rearNode = rearNode.next
	}

	for {
		//头节点就是需要删除的节点
		if tailNode == delNode {
			newLink = tailNode.next
			rearNode.next = newLink
			break
		}
		//中间节点是需要删除的节点
		if tailNode.next == delNode {
			tailNode.next = delNode.next
			break
		}
		//尾节点是需要删除的节点
		if tailNode.next == ln {
			if tailNode == delNode {
				//找到倒数第二个节点
				descSecondNode := ln
				for {
					if descSecondNode.next == tailNode {
						break
					}
					descSecondNode = descSecondNode.next
				}
				descSecondNode.next = ln
			} else {
				return nil, fmt.Errorf("未找到要删除的节点")
			}
			break
		}
		tailNode = tailNode.next
	}
	return
}

//
// showLink
// @Description: 遍历链表
// @Author: maxwell.ke
// @time 2022-10-26 15:54:18
// @receiver ln
//
func (ln *LinkNode) showLink() {
	tail := ln
	if tail.next == nil {
		fmt.Println("链表为空")
		return
	}
	for {
		fmt.Printf("%s-->", tail.value)
		if tail.next == ln {
			break
		}
		tail = tail.next
	}
	fmt.Println()
}
