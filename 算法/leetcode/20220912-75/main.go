package main

import "fmt"

func main() {
	arr := []int{2, 0, 2, 1, 1, 0}
	sortColors(arr)
	fmt.Println(arr)
}

//func sortColors(nums []int) {
//	if len(nums) <= 1 {
//		return
//	}
//	buckets := make([][]int, 3)
//	for k, v := range nums {
//		buckets[v] = append(buckets[v], k)
//	}
//	newArr := make([]int, 0, len(nums))
//	for _, indexArr := range buckets {
//		for _, index := range indexArr {
//			newArr = append(newArr, nums[index])
//		}
//	}
//	nums = newArr
//	fmt.Println(nums)
//}

func sortColors(nums []int) {
	count0 := swapColors(nums, 0) // 把 0 排到前面
	swapColors(nums[count0:], 1)  // nums[:count0] 全部是 0 了，对剩下的 nums[count0:] 把 1 排到前面
}

func swapColors(colors []int, target int) (countTarget int) {
	for i, c := range colors {
		if c == target {
			colors[i], colors[countTarget] = colors[countTarget], colors[i]
			countTarget++
		}
	}
	return
}
