package main

import (
	"fmt"
	"sort"
)

func main() {
	//nums := []int{3, 6, 9, 1}
	nums := []int{10}
	fmt.Println(maximumGap(nums))
}

func maximumGap(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}
	sort.Ints(nums)
	var bigGap int
	for i := 0; i < n-1; i++ {
		if nums[i+1]-nums[i] > bigGap {
			bigGap = nums[i+1] - nums[i]
		}
	}
	return bigGap
}
