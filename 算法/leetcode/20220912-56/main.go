package main

import (
	"fmt"
	"sort"
)

func main() {
	//intervals := [][]int{[]int{1, 3}, []int{2, 9}, []int{8, 10}, []int{15, 18}, []int{2, 6}}
	intervals := [][]int{[]int{2, 3}, []int{4, 5}, []int{6, 7}, []int{8, 9}, []int{1, 10}}
	arr := merge(intervals)
	fmt.Println(arr)
}

func merge(intervals [][]int) [][]int {
	n := len(intervals)
	if n <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1]
	})
	fmt.Println(intervals)

	newIntervals := make([][]int, 0, len(intervals))

	for i := 0; i < n; i++ {
		L, R := intervals[i][0], intervals[i][1]
		if len(newIntervals) == 0 || newIntervals[len(newIntervals)-1][1] < L {
			newIntervals = append(newIntervals, []int{L, R})
		} else {
			newIntervals[len(newIntervals)-1][1] = max(newIntervals[len(newIntervals)-1][1], R)
		}
	}

	return newIntervals
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
