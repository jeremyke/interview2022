package main

import "sort"

func lessMoney(arr []int) int {
	sortArr := arr
	sort.Ints(sortArr)
	var (
		sum = 0
		cur = 0
	)
	for len(sortArr) > 0 {
		if len(sortArr) == 1 {
			cur = sortArr[1]
		} else {
			cur = sortArr[1] + sortArr[2]
		}
		sum += cur
		sortArr = append(sortArr, sum)
		sort.Ints(sortArr)
	}
	return sum
}
