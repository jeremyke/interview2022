package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}
	fmt.Println(threeSum(nums))
}

//三数字之和

//暴力解法
func threeSum1(nums []int) [][]int {
	sort.Ints(nums)
	resArr := make([][]int, 0, 1)
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := len(nums) - 1; k > j; k-- {
				if nums[i]+nums[j]+nums[k] == 0 {
					//判断是否有重复的
					if !isRepeat(resArr, nums[i], nums[j], nums[k]) {
						resArr = append(resArr, []int{nums[i], nums[j], nums[k]})
					}
					break
				}
			}
		}
	}
	return resArr
}

func isRepeat(resArr [][]int, i, j, k int) bool {
	if len(resArr) > 0 {
		for m := 0; m < len(resArr); m++ {
			var iRes, jRes, kRes bool
			if i == 0 && j == 0 && k == 0 {
				if resArr[m][0] == 0 && resArr[m][1] == 0 && resArr[m][2] == 0 {
					iRes, jRes, kRes = true, true, true
				}
			} else {
				for n := 0; n < 3; n++ {
					if resArr[m][n] == i {
						iRes = true
					}
					if resArr[m][n] == j {
						jRes = true
					}
					if resArr[m][n] == k {
						kRes = true
					}
				}
			}
			if iRes && jRes && kRes {
				return true
			}
		}
		return false
	} else {
		return false
	}
}

//排序+双指针法
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	resArr := make([][]int, 0, 1)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1
		k := len(nums) - 1
		for j < k {
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			if k < len(nums)-1 && nums[k] == nums[k+1] {
				k--
				continue
			}
			if nums[i]+nums[j]+nums[k] == 0 {
				resArr = append(resArr, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			}
			if nums[i]+nums[j]+nums[k] > 0 {
				k--
			}
			if nums[i]+nums[j]+nums[k] < 0 {
				j++
			}
		}
	}
	return resArr
}
