package main

import (
	"fmt"
	"sort"
)

//给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。

func main() {
	//nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	//nums := []int{3, 0, 1}
	//nums := []int{0, 1}
	nums := []int{1}
	fmt.Println(missingNumber(nums))
}

func missingNumber(nums []int) int {
	sort.Ints(nums)
	if nums[0] == 0 {
		nums = nums[1:]
		nums = append(nums, 0)
	}
	var missNum int
	for k, v := range nums {
		if k != v-1 {
			missNum = k + 1
			break
		}
	}
	return missNum
}
