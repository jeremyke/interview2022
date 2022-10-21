package main

//赢的方法
func win1(arr []int, i, j int) int {
	if len(arr) == 0 {
		return 0
	}
	return max(f(arr, 0, len(arr)-1), s(arr, 0, len(arr)-1))
}

//先手函数
func f(arr []int, i, j int) int {
	if i == j {
		return arr[i]
	}
	return max(arr[i]+s(arr, i+1, j), arr[j]+s(arr, i, j-1))
}

//后手函数
func s(arr []int, i, j int) int {
	if i == j {
		return 0
	}
	return min(f(arr, i+1, j), f(arr, i, j-1))
}

func max(i, j int) int {
	big := i
	if j > i {
		big = j
	}
	return big
}

func min(i, j int) int {
	minNum := i
	if j < i {
		minNum = j
	}
	return minNum
}
