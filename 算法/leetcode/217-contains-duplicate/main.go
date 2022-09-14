package main

import "fmt"

//给你一个整数数组 nums 。如果任一值在数组中出现 至少两次 ，返回 true ；如果数组中每个元素互不相同，返回 false 。

func main() {
	//arr := []int{1, 2, 3, 1}
	//arr := []int{1, 2, 3, 4}
	arr := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(containsDuplicate(arr))
}

func containsDuplicate(nums []int) bool {
	aMap := make(map[int]int, len(nums))
	flag := false
	for _, v := range nums {
		if _, ok := aMap[v]; ok {
			flag = true
			break
		} else {
			aMap[v]++
		}
	}
	return flag
}
