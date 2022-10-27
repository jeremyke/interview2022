package main

import (
	"errors"
	"fmt"
	"math"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type Stack struct {
	MaxTop int      //栈最大可以存放的数的个数
	Top    int      //表示栈顶的索引id，初始值为-1，最大值为MaxTop-1
	Arr    [7]*Node //数组模拟栈
}

func (s *Stack) Push(pushNode *Node) error {
	if s.Top == s.MaxTop-1 {
		return errors.New("栈满了")
	}

	s.Top++
	s.Arr[s.Top] = pushNode

	return nil
}

func (s *Stack) Pop() (popNode *Node, err error) {
	if s.Top == -1 {
		return nil, errors.New("空栈")
	}
	popNode = s.Arr[s.Top]
	s.Arr[s.Top] = nil
	s.Top--
	return popNode, nil
}

func (s *Stack) List() {
	if s.Top == -1 {
		fmt.Println("空栈")
	}
	for i := s.Top; i >= 0; i-- {
		fmt.Printf("Arr[%v]=%v\n", i, s.Arr[i].Value)
	}
}

func (s *Stack) IsEmpty() bool {
	flag := false
	if s.Top == -1 {
		flag = true
	}
	return flag
}

func (s *Stack) IsFull() bool {
	flag := false
	if s.Top == s.MaxTop-1 {
		flag = true
	}
	return flag
}

type Queue struct {
	buff    []*Node //队列的的数据存储在数组上
	maxsize int     //队列最大容量
	front   int     //队列头索引，不包括自己（队列头索引值-1）
	rear    int     //队列尾索引
}

//
// Push
// @Description: 压入队列
// @Author: maxwell.ke
// @time 2022-10-25 22:58:58
// @receiver q
// @param n
// @return error
//
func (q *Queue) Push(pushNode *Node) error {
	if q.rear == q.maxsize-1 {
		if q.front == -1 { //头尾都到头了
			return fmt.Errorf("队列已满，PUSH失败")
		} else { ////尾都到头了，头有空，则重置front
			q.front = -1
			q.rear = len(q.buff) - 1
		}
	}
	q.rear++
	q.buff = append(q.buff, pushNode)
	return nil
}

//
// Pop
// @Description: 出队列
// @Author: maxwell.ke
// @time 2022-10-25 23:14:20
// @receiver q
// @return n
// @return err
//
func (q *Queue) Pop() (popNode *Node, err error) {
	if len(q.buff) == 0 {
		return nil, fmt.Errorf("空队列,POP失败")
	}
	popNode = q.buff[0]
	q.buff = q.buff[1:]
	q.front++
	return popNode, nil
}

//
// List
// @Description: 队列遍历
// @Author: maxwell.ke
// @time 2022-10-25 23:13:10
// @receiver q
// @return error
//
func (q *Queue) List() error {
	if len(q.buff) == 0 {
		return fmt.Errorf("空队列")
	}
	for i := 0; i < q.maxsize; i++ {
		if i > q.front && i <= q.rear {
			fmt.Println(q.buff[i-q.front-1].Value)
		} else {
			fmt.Println("nil")
		}
	}
	fmt.Println()
	return nil
}

func main() {
	root := &Node{Value: 1}
	left := &Node{Value: 3}
	right := &Node{Value: 4}
	root.Left = left
	root.Right = right
	left1 := &Node{Value: 5}
	right1 := &Node{Value: 6}
	left2 := &Node{Value: 7}
	right2 := &Node{Value: 8}
	left3 := &Node{Value: 9}
	left.Left = left1
	left.Right = right1
	right.Left = left2
	right.Right = right2
	left1.Left = left3
	//preShowTree(root)
	//midShowTree(root)
	afterShowTree(root)
	//wideShowTree(root)
	fmt.Println(isBST(root))
	fmt.Println(isFBT(root))
	fmt.Println(isCBT(root))
	fmt.Println(isALV(root))
}

func preShowTree(head *Node) {
	if head != nil {
		nodeStack := &Stack{MaxTop: 7, Top: -1}
		nodeStack.Push(head)
		for nodeStack.Top < nodeStack.MaxTop-1 && nodeStack.Top > -1 {
			tmpNode, _ := nodeStack.Pop()
			if tmpNode != nil {
				fmt.Println(tmpNode.Value)
			}
			if tmpNode.Right != nil {
				nodeStack.Push(tmpNode.Right)
			}
			if tmpNode.Left != nil {
				nodeStack.Push(tmpNode.Left)
			}
		}
	}
}

//
// midShowTree
// @Description: 中序遍历
//1. 申请1个栈
//2. 把每颗子树的整树所有左节点入栈
//3. 依次弹出的过程中，打印该节点，且对该节点的右节点，循环1，2,
//4. 循环执行1->2->3
//5. 打印收集栈中数据
// @Author: maxwell.ke
// @time 2022-10-26 21:59:31
// @param head
//
func midShowTree(head *Node) {
	if head != nil {
		nodeStack := &Stack{MaxTop: 7, Top: -1}
		for !nodeStack.IsEmpty() || head != nil {
			if head != nil { //所有左节点压入栈
				nodeStack.Push(head)
				head = head.Left
			} else { //弹出栈的数据，打印节点，且对右子节点循环if的操作
				popNode, _ := nodeStack.Pop()
				fmt.Println(popNode.Value)
				head = popNode.Right
			}
		}
	}
}

//
// afterShowTree
// @Description:非递归后序遍历
//1. 申请两个栈（普通栈和收集栈）
//2. 把当前节点压入普通栈中
//2. 从普通栈中弹出一个节点cur
//3. 把该节点压入收集栈中
//4. 先左后右（如果有），把该节点的子节点压入普通栈中
//5. 循环执行1->2->3
//6. 打印收集栈中数据
// @Author: maxwell.ke
// @time 2022-10-26 22:14:20
// @param head
//
func afterShowTree(head *Node) {
	if head != nil {
		nodeStack := &Stack{MaxTop: 7, Top: -1}
		showStack := &Stack{MaxTop: 7, Top: -1}
		nodeStack.Push(head)
		for !nodeStack.IsEmpty() {
			popNode, _ := nodeStack.Pop()
			showStack.Push(popNode) //当前节点压入收集栈
			//左子节点压入普通栈
			if popNode.Left != nil {
				nodeStack.Push(popNode.Left)
			}
			//右子节点压入栈
			if popNode.Right != nil {
				nodeStack.Push(popNode.Right)
			}
		}
		//便利收集栈数据
		showStack.List()
	}
}

//
// wideShowTree
// @Description: 宽度优先遍历
// @Author: maxwell.ke
// @time 2022-10-26 22:50:57
// @param head
//
func wideShowTree(head *Node) {
	nodeQueue := &Queue{buff: make([]*Node, 0, 7), maxsize: 7, front: -1, rear: -1}
	nodeQueue.Push(head)
	for len(nodeQueue.buff) < nodeQueue.maxsize && len(nodeQueue.buff) > 0 {
		cur, _ := nodeQueue.Pop()
		fmt.Println(cur.Value)
		if cur.Left != nil {
			nodeQueue.Push(cur.Left)
		}
		if cur.Right != nil {
			nodeQueue.Push(cur.Right)
		}
	}
}

var preValue int = math.MinInt64

//
// isSBT
// @Description: 搜索二叉树
// @Author: maxwell.ke
// @time 2022-10-26 23:50:57
// @param head
// @return bool
//
func isBST(head *Node) bool {
	if head == nil {
		return true
	}
	leftBST := isBST(head.Left)
	if !leftBST {
		return false
	}
	//打印事件改为了比较事件
	if head.Value <= preValue {
		return false
	} else {
		preValue = head.Value
	}

	rightBST := isBST(head.Right)
	return rightBST
}

//
// isFBT
// @Description: 满二叉树
// @Author: maxwell.ke
// @time 2022-10-26 23:51:33
// @param head
// @return bool
//
func isFBT(head *Node) bool {
	if head == nil {
		return true
	}
	leftFBT := isFBT(head.Left)
	if !leftFBT {
		return false
	}
	if (head.Left != nil && head.Right == nil) || (head.Left == nil && head.Right != nil) {
		return false
	}
	rightFBT := isFBT(head.Right)
	return rightFBT
}

//
// isCBT
// @Description: 完全二叉树
// @Author: maxwell.ke
// @time 2022-10-27 00:09:42
// @param head
// @return bool
//
func isCBT(head *Node) bool {
	if head == nil {
		return true
	}
	var (
		leaf = false //是否遇到左右子节点不齐全的情况
		l    *Node
		r    *Node
	)
	nodeQueue := &Queue{buff: make([]*Node, 0, 8), maxsize: 8, front: -1, rear: -1}
	nodeQueue.Push(head)
	for len(nodeQueue.buff) < nodeQueue.maxsize && len(nodeQueue.buff) > 0 {
		popNode, _ := nodeQueue.Pop()
		l = popNode.Left
		r = popNode.Right

		//如果遇到了不双全的节点后，又发现当前节点有子节点 或者 有右节点但无左节点
		if (leaf && (l != nil || r != nil)) || (l == nil && r != nil) {
			return false
		}

		if l != nil {
			nodeQueue.Push(l)
		}
		if r != nil {
			nodeQueue.Push(r)
		}
		if l == nil || r == nil { //左右节点不双全，则后续为叶子节点
			leaf = true
		}
	}
	return true
}

//
// isALV
// @Description:平衡二叉树
// @Author: maxwell.ke
// @time 2022-10-27 00:18:08
// @param head
// @return bool
//
func isALV(head *Node) bool {
	return process(head).IsBalanced
}

type ReturnType struct {
	IsBalanced bool
	Height     int
}

func process(x *Node) *ReturnType {
	if x == nil {
		return &ReturnType{IsBalanced: true, Height: 0}
	}
	leftData := process(x.Left)
	rightData := process(x.Right)
	height := max(leftData.Height, rightData.Height) + 1 //当前节点的高度
	isBalanced := leftData.IsBalanced && rightData.IsBalanced && math.Abs(float64(leftData.Height-rightData.Height)) <= 1
	return &ReturnType{IsBalanced: isBalanced, Height: height}
}

func max(a, b int) int {
	maxNum := a
	if a < b {
		maxNum = b
	}
	return maxNum
}
