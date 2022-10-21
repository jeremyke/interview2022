package main

import "fmt"

//移除栈底元素，上面的元素往下移动一位
func f(zhan []int) (bottom int, newStack []int) {
	if len(zhan) == 0 {
		return 0, zhan
	}
	bottom = zhan[0]
	newStack = zhan[1:]
	return bottom, newStack
}

func reverse(zhan []int) {
	if len(zhan) == 0 {
		return
	}
	i, newStack := f(zhan)
	reverse(newStack)
	newStack = append(newStack, i)
}

func main() {
	zhan := []int{3, 2, 1}
	reverse(zhan)
	fmt.Println(zhan)
}
