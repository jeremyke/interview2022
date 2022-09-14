package main

import "fmt"

func main() {
	nums, k := []int{3, 2, 1, 5, 6, 4}, 2
	//nums, k := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4
	fmt.Println(findKthLargest(nums, k))
}

func findKthLargest(nums []int, k int) int {
	//先对数组进行排序O(N)
	QuickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
	//取出倒数第n个数
	return nums[len(nums)-k]

}

//func partition(arr []int, l, r int) []int {
//	less := l - 1 //<区右边界
//	more := r     //>区左边界
//	for l < more {
//		if arr[l] < arr[r] {
//			less++
//			swap(arr, less, l)
//			less++
//			l++
//		} else if arr[l] > arr[r] {
//			more--
//			swap(arr, more, l)
//		} else {
//			l++
//		}
//	}
//	swap(arr, more, r)
//	return []int{less + 1, more}
//}

func QuickSort(array []int, left, right int) {
	l := left
	r := right

	pivot := array[(left+right)/2]
	//for循环的目标是将比pivot小的数放在左边，比pivot大的数放在右边
	for l < r {
		//从pivot左边找到一个大于等于pivot的值
		for array[l] < pivot {
			l++
		}
		//从pivot右边找到小于等于pivot的值
		for array[r] > pivot {
			r--
		}
		//表明本次分解任务完成
		if l >= r {
			break
		}
		//交换
		array[l], array[r] = array[r], array[l]
		//优化
		if array[l] == pivot {
			r--
		}
		if array[r] == pivot {
			l++
		}
		//	fmt.Printf("%v\n", *array)
	}
	//可注释以下代码做测试
	//如果l==r，在移动下
	if l == r {
		l++
		r--
	}
	//向左递归
	if left < r {
		QuickSort(array, left, r)
	}
	//向右递归
	if right > l {
		QuickSort(array, l, right)
	}
}

func swap(arr []int, a, b int) {
	arr[a] = arr[a] ^ arr[b]
	arr[b] = arr[a] ^ arr[b]
	arr[a] = arr[a] ^ arr[b]
	return
}
