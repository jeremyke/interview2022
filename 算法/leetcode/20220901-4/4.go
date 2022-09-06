package main

import "fmt"

func main() {
	n1 := []int{1}
	n2 := []int{2, 3, 4}
	fmt.Println(findMedianSortedArrays(n1, n2))
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var leftIndex, rightIndex int
	var res float64
	combinSlice := make([]int, 0, 1)
	if (len(nums1)+len(nums2))%2 == 0 {
		leftIndex, rightIndex = (len(nums1)+len(nums2))/2-1, (len(nums1)+len(nums2))/2
	} else {
		leftIndex, rightIndex = (len(nums1)+len(nums2)+1)/2-1, (len(nums1)+len(nums2)+1)/2-1
	}
	if len(nums1) == 0 && len(nums2) != 0 {
		res = float64(nums2[leftIndex]+nums2[rightIndex]) / 2
	} else if len(nums1) != 0 && len(nums2) == 0 {
		res = float64(nums1[leftIndex]+nums1[rightIndex]) / 2
	} else if len(nums1) == 0 && len(nums2) == 0 {
		res = 0
	} else {
		var firstSlice, secondSlice []int
		if nums1[0] <= nums2[0] {
			firstSlice = nums1
			secondSlice = nums2
		} else {
			firstSlice = nums2
			secondSlice = nums1
		}
		combinSlice = append(combinSlice, firstSlice[0])
		var k2 int
		for k1 := 0; k1 < len(firstSlice); {
			for k2 < len(secondSlice) {
				if k1 == len(firstSlice)-1 {
					combinSlice = append(combinSlice, secondSlice[k2])
					if len(combinSlice) >= rightIndex+1 {
						break
					}
					k2++
				} else {
					if secondSlice[k2] <= firstSlice[k1+1] {
						combinSlice = append(combinSlice, secondSlice[k2])
						if len(combinSlice) >= rightIndex+1 {
							break
						}
						k2++
					} else {
						k1++
						combinSlice = append(combinSlice, firstSlice[k1])
						if len(combinSlice) >= rightIndex+1 {
							break
						}
					}
				}
			}
			if len(combinSlice) >= rightIndex+1 {
				break
			}
			k1++
			combinSlice = append(combinSlice, firstSlice[k1])
		}
		res = float64(combinSlice[leftIndex]+combinSlice[rightIndex]) / 2
	}
	return res
}
