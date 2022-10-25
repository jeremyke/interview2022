package main

import "fmt"

type Queue struct {
	buff    []int //队列的的数据存储在数组上
	maxsize int   //队列最大容量
	front   int   //队列头索引，不包括自己（队列头索引值-1）
	rear    int   //队列尾索引
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
func (q *Queue) Push(n int) error {
	if q.rear == q.maxsize-1 {
		if q.front == -1 { //头尾都到头了
			return fmt.Errorf("队列已满，PUSH失败")
		} else { ////尾都到头了，头有空，则重置front
			q.front = -1
			q.rear = len(q.buff) - 1
		}
	}
	q.rear++
	q.buff = append(q.buff, n)
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
func (q *Queue) Pop() (n int, err error) {
	if len(q.buff) == 0 {
		return 0, fmt.Errorf("空队列,POP失败")
	}
	n = q.buff[0]
	q.buff = q.buff[1:]
	q.front++
	return n, nil
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
			fmt.Println(q.buff[i-q.front-1])
		} else {
			fmt.Println("nil")
		}
	}
	fmt.Println()
	return nil
}

func main() {
	var aQueue = Queue{
		buff:    make([]int, 0, 5),
		maxsize: 5,
		front:   -1,
		rear:    -1,
	}
	aQueue.Push(3)
	aQueue.Push(6)
	aQueue.Push(5)
	aQueue.List()
	aQueue.Pop()
	aQueue.Pop()
	aQueue.List()
	aQueue.Push(1)
	aQueue.Push(2)
	aQueue.Push(3)
	aQueue.List()
}
