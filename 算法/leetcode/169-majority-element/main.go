package main

import (
	"fmt"
	"sort"
)

//给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
//
//你可以假设数组是非空的，并且给定的数组总是存在多数元素。
//

func main() {
	//arr := []int{3, 2, 3}
	arr := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(arr))
}

func majorityElement(nums []int) int {
	sort.Ints(nums)
	var (
		n         = len(nums)
		startId   int
		tmpNum    int
		maxNumber int
	)
	if n == 1 {
		return nums[0]
	}

	for i := 0; i < n; i++ {
		if i == 0 {
			tmpNum = nums[0]
			startId = 0
		} else {
			if i-startId+1 > n/2 {
				maxNumber = tmpNum
			}
			if nums[i] != nums[i-1] {
				tmpNum = nums[i]
				startId = i
			}
		}
	}
	return maxNumber

}
