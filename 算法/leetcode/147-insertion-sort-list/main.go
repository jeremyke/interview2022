package main

import "fmt"

//给定单个链表的头 head ，使用 插入排序 对链表进行排序，并返回 排序后链表的头 。
//
//插入排序 算法的步骤:
//
//插入排序是迭代的，每次只移动一个元素，直到所有元素可以形成一个有序的输出列表。
//每次迭代中，插入排序只从输入数据中移除一个待排序的元素，找到它在序列中适当的位置，并将其插入。
//重复直到所有输入数据插入完为止。
//下面是插入排序算法的一个图形示例。部分排序的列表(黑色)最初只包含列表中的第一个元素。每次迭代时，从输入数据中删除一个元素(红色)，并就地插入已排序的列表中。
//
//对链表进行插入排序。
//

func main() {
	headArr := []int{4, 2, 1, 3}
	n := len(headArr)
	//对数组进行插入排序
	for i := 1; i < n; i++ {
		for j := i; j > 0 && headArr[j] < headArr[j-1]; j-- {
			swap(headArr, j, j-1)
		}
	}
	fmt.Println(headArr)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	tmpNode := head
	headArr := make([]int, 0, 1)
	for {
		if tmpNode == nil {
			break
		}
		headArr = append(headArr, tmpNode.Val)
		tmpNode = tmpNode.Next
	}
	n := len(headArr)
	if n <= 1 {
		return head
	}
	//对数组进行插入排序
	for i := 1; i < n; i++ {
		for j := i; j > 0 && headArr[j] < headArr[j-1]; j-- {
			swap(headArr, j, j-1)
		}
	}
	//便历数组生成node
	var newList *ListNode
	for k := n - 1; k >= 0; k-- {
		tempNode := ListNode{Val: headArr[k]}
		if newList == nil {
			newList = &tempNode
		} else {
			tempNode.Next = newList
			newList = &tempNode
		}
	}

	return newList
}

func swap(arr []int, a, b int) {
	arr[a] = arr[a] ^ arr[b]
	arr[b] = arr[a] ^ arr[b]
	arr[a] = arr[a] ^ arr[b]
	return
}
