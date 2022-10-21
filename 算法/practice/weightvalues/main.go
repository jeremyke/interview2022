package main

import (
	"fmt"
)

func getTotalValue(weights []int, values []int, bag int) int {
	if len(weights) == 0 || len(values) == 0 || bag == 0 {
		return 0
	}
	return process(weights, values, 0, 0, 0, bag)
}

func process(weights []int, values []int, i, totalWeight, totalValue, bag int) int {
	if totalWeight > bag {
		return 0
	}
	if len(weights) == i {
		return totalValue
	}
	return max(process(weights, values, i+1, totalWeight, totalValue, bag), process(weights, values, i+1, totalWeight+weights[i], totalValue+values[i], bag))
}

func max(i, j int) int {
	big := i
	if j > i {
		big = j
	}
	return big
}

func main() {
	weights := []int{1, 2, 3, 4}
	values := []int{3, 1, 7, 8}
	bag := 3
	res := getTotalValue(weights, values, bag)
	fmt.Println(res)
}
