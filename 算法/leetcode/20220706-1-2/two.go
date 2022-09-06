package main

//给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//链表问题需要注意 头部和尾部
func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		res   *ListNode
		tail  *ListNode
		extra int
	)
	for l1 != nil || l2 != nil {
		var l1Val, l2Val int
		if l1 == nil {
			l1Val = 0
		} else {
			l1Val = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			l2Val = 0
		} else {
			l2Val = l2.Val
			l2 = l2.Next
		}
		resVal := (l1Val + l2Val + extra) % 10
		extra = (l1Val + l2Val + extra) / 10
		resTmp := &ListNode{
			Val: resVal,
		}
		if res != nil {
			tail.Next = resTmp
			tail = tail.Next
		} else {
			res = resTmp
			tail = res
		}
	}
	if extra > 0 {
		tail.Next = &ListNode{Val: extra}
	}
	return res
}
