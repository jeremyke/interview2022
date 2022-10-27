package main

import "fmt"

type joshefu struct {
	no   int
	next *joshefu
}

/**
 * @Description: 创建一个链表
 * @Param:num:链表的长度
 * @Return:返回环形链表的头结点
 */
func addLink(num int) *joshefu {
	if num <= 0 {
		return nil
	}
	head := &joshefu{}
	p := head
	for i := 1; i <= num; i++ {
		p.no = i
		if i == num {
			p.next = head
		} else {
			p.next = &joshefu{}
		}
		p = p.next
	}
	return head
}

func showLink(head *joshefu) {
	p := head
	for {
		if p.next == nil {
			fmt.Println("链表为空")
			break
		}
		fmt.Printf("%d -->", p.no)
		if p.next == head {
			break
		}
		p = p.next
	}
	fmt.Println()
}

func delLink(head *joshefu, start int, step int) {
	if head.next == nil {
		fmt.Println("链表为空")
		return
	}
	//寻找尾节点
	rearNode := head
	for {
		if rearNode.next == head {
			break
		}
		rearNode = rearNode.next
	}

	//让head移动到start节点（start-1）
	for i := 1; i <= start-1; i++ {
		head = head.next
		rearNode = rearNode.next
	}

	//创建一个数组
	var delArr []int

	fmt.Println(head.no)
	//找到删除节点并删除
	for {
		for j := 1; j <= step-1; j++ {
			head = head.next
			rearNode = rearNode.next
		}
		delArr = append(delArr, head.no)
		head = head.next
		rearNode.next = head
		if head == rearNode {
			delArr = append(delArr, head.no)
			break
		}

	}
	//展示出列顺序
	fmt.Println(delArr)
}

func main() {
	linkHead := addLink(10)
	showLink(linkHead)
	delLink(linkHead, 3, 4)
}
