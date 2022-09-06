package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	//nums := []int{-1, 2, 1, -4}
	//nums := []int{1, 1, -1}
	//nums := []int{4, 0, 5, -5, 3, 3, 0, -4, -5}
	nums := []int{-13, 592, -501, 770, -952, -675, 322, -829, -246, 657, 608, 485, -112, 967, -30, 182, -969, 559, -286, -64, 24, 365, -158, 701, 535, -429, -217, 28, 948, -114, -536, -711, 693, 23, -958, -283, -700, -672, 311, 314, -712, -594, -351, 658, 747, 949, 70, 888, 166, 495, 244, -380, -654, 454, -281, -811, -168, -839, -106, 877, -216, 523, -234, -8, 289, -175, 920, -237, -791, -976, -509, -4, -3, 298, -190, 194, -328, 265, 150, 210, 285, -176, -646, -465, -97, -107, 668, 892, 612, -54, -272, -910, 557, -212, -930, -198, 38, -365, -729, -410, 932, 4, -565, -329, -456, 224, 443, -529, -428, -294, 191, 229, 112, -867, -163, -979, 236, -227, -388, -209, 984, 188, -549, 970, 951, -119, -146, 801, -554, 564, -769, 334, -819, -356, -724, -219, 527, -405, -27, -759, 722, -774, 758, 394, 146, 517, 870, -208, 742, -782, 336, -364, -558, -417, 663, -914, 536, 293, -818, 847, -322, 408, 876, -823, 827, 167}
	fmt.Println(threeSumClosest1(nums, 7175))
}

//最接近的三数之和
func threeSumClosest1(nums []int, target int) int {
	sort.Ints(nums)
	//var he = nums[len(nums)-1] + nums[len(nums)-1] + nums[len(nums)-1]
	var he = math.MaxInt32
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := len(nums) - 1
		for j < k {
			tmpHe := nums[i] + nums[j] + nums[k]
			if tmpHe == target {
				he = target
				return he
			}
			if abs(tmpHe-target) < abs(he-target) {
				he = tmpHe
			}
			if tmpHe > target {
				k0 := k - 1
				// 移动到下一个不相等的元素
				for j < k0 && nums[k0] == nums[k] {
					k0--
				}
				k = k0
			} else {
				j0 := j + 1
				// 移动到下一个不相等的元素
				for j0 < k && nums[j0] == nums[j] {
					j0++
				}
				j = j0
			}
		}
	}
	return he
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var (
		n    = len(nums)
		best = math.MaxInt32
	)

	// 根据差值的绝对值来更新答案
	update := func(cur int) {
		if abs(cur-target) < abs(best-target) {
			best = cur
		}
	}

	// 枚举 a
	for i := 0; i < n; i++ {
		// 保证和上一次枚举的元素不相等
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 使用双指针枚举 b 和 c
		j, k := i+1, n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			// 如果和为 target 直接返回答案
			if sum == target {
				return target
			}
			update(sum)
			if sum > target {
				// 如果和大于 target，移动 c 对应的指针
				k0 := k - 1
				// 移动到下一个不相等的元素
				for j < k0 && nums[k0] == nums[k] {
					k0--
				}
				k = k0
			} else {
				// 如果和小于 target，移动 b 对应的指针
				j0 := j + 1
				// 移动到下一个不相等的元素
				for j0 < k && nums[j0] == nums[j] {
					j0++
				}
				j = j0
			}
		}
	}
	return best
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
