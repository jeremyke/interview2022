package main

import "fmt"

func main() {
	ss := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(ss))
}

//暴力解法
func maxArea1(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	var area int
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			h := height[i]
			if height[i] > height[j] {
				h = height[j]
			}
			l := j - i
			tmpAre := h * l
			//fmt.Println(i, j, tmpAre)
			if area < tmpAre {
				area = tmpAre
			}
		}
	}
	return area
}

//在每个状态下，无论长板或短板向中间收窄一格，都会导致水槽 底边宽度 -1−1​ 变短：
//
//若向内 移动短板 ，水槽的短板 min(h[i], h[j])min(h[i],h[j]) 可能变大，因此下个水槽的面积 可能增大 。
//若向内 移动长板 ，水槽的短板 min(h[i], h[j])min(h[i],h[j])​ 不变或变小，因此下个水槽的面积 一定变小 。
//因此，初始化双指针分列水槽左右两端，循环每轮将短板向内移动一格，并更新面积最大值，直到两指针相遇时跳出；即可获得最大面积。
//
func maxArea(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	left := 0
	right := len(height) - 1
	var area int
	for left <= right && right >= left {
		duan := height[left]
		flag := "l"
		if height[left] > height[right] {
			duan = height[right]
			flag = "r"
		}
		tmpArea := duan * (right - left)
		if tmpArea > area {
			area = tmpArea
		}
		if flag == "l" {
			left++
		} else {
			right--
		}
	}
	return area
}
