package main

import "fmt"

func main() {
	arr := []int{3, 1, 4, 6, 8, 2, 5, 13, 76, 11}
	fmt.Println(process(arr, 0, 9))
}

func process(arr []int, i, j int) []int {
	if i == j {
		return arr
	}
	mid := i + (j-i)>>1
	fmt.Println(mid)
	process(arr, i, mid)
	process(arr, mid+1, j)
	merge(arr, i, mid, j)
	return arr
}

func merge(arr []int, i, mid, j int) {
	newArr := make([]int, j-i+1)
	a := 0
	p1 := i
	p2 := mid + 1
	for p1 <= mid && p2 <= j {
		if arr[p1] > arr[p2] {
			newArr[a+1] = arr[p2+1]
		} else {
			newArr[a+1] = arr[p1+1]
		}
	}
	for p1 < mid {
		newArr[a+1] = arr[p1+1]
	}
	for p2 < j {
		newArr[a+1] = arr[p2+1]
	}
	for m := 0; m < len(newArr); m++ {
		arr[i+m] = newArr[m]
	}
}
