package main

import "fmt"

//给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
//
//请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
//
//注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。
//

func main() {
	//nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
	//nums1, m, nums2, n := []int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3
	//nums1, m, nums2, n := []int{1}, 1, []int{}, 0
	nums1, m, nums2, n := []int{0}, 0, []int{1}, 1
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}
	for j := 0; j < m+n-1; j++ {
		for k := j + 1; k < m+n; k++ {
			if nums1[j] > nums1[k] {
				swap(nums1, j, k)
			}
		}
	}
}
func swap(arr []int, a, b int) {
	arr[a] = arr[a] ^ arr[b]
	arr[b] = arr[a] ^ arr[b]
	arr[a] = arr[a] ^ arr[b]
	return
}
