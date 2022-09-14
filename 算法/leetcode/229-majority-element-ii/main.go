package main

import "fmt"

//给定一个大小为 n 的整数数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。

func main() {
	//nums := []int{3, 2, 3}
	nums := []int{1, 2}
	fmt.Println(majorityElement(nums))
}

func majorityElement(nums []int) []int {
	var (
		lens   = len(nums)
		bucket = make(map[int][]int, lens)
		res    = make([]int, 0, 1)
	)

	for k, v := range nums {
		bucket[v] = append(bucket[v], k)
	}
	for k1, v1 := range bucket {
		if len(v1) > lens/3 {
			res = append(res, k1)
		}
	}

	return res
}
