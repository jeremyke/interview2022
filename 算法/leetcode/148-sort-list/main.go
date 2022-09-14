package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	tmpNode := head
	//链表转为数组
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

func swap(arr []int, i, j int) {
	arr[i] = arr[i] ^ arr[j]
	arr[j] = arr[i] ^ arr[j]
	arr[i] = arr[i] ^ arr[j]
}
