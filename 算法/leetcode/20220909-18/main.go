package main

import "sort"

//给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：
//
//0 <= a, b, c, d < n
//a、b、c 和 d 互不相同
//nums[a] + nums[b] + nums[c] + nums[d] == target
//你可以按 任意顺序 返回答案 。
//

//思路：排序+双指针降低一个0(N)

func main() {
	arr, target := []int{1, 0, -1, 0, -2, 2}, 0
	println(fourSum(arr, target))
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var (
		n   = len(nums)
		res = make([][]int, 0, 1)
	)

	//循环确定第一个数
	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		//(1)第i个数和前面第i-1个数相同时，跳过
		//(2)第i个数和最大的三个数之和小于target，直接跳出循环
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		//循环确定第二个数
		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			//(1)第j个数和前面第j-1个数相同时，跳过
			//(2)第i个数和第j个数，和最大的两个数之和小于target，直接跳出循环
			if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}
			//双指针确定最后两个数
			for left, right := j+1, n-1; left < right; {
				if sum := nums[i] + nums[j] + nums[left] + nums[right]; sum == target {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					//跳过left==left前一个数的数
					for left++; left < right && nums[left] == nums[left-1]; left++ {
					}
					//跳过right==right前一个数的数
					for right--; left < right && nums[right] == nums[right+1]; right-- {
					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return res
}
