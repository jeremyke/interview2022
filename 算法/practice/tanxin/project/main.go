package main

import "sort"

type node struct {
	c int
	p int
}

type nodeCArr []*node

func (array nodeCArr) Len() int {
	return len(array)
}

func (array nodeCArr) Less(i, j int) bool {
	return array[i].c < array[j].c //从小到大， 若为大于号，则从大到小
}

func (array nodeCArr) Swap(i, j int) {
	array[i], array[j] = array[j], array[i]
}

type nodePArr []*node

func (array nodePArr) Len() int {
	return len(array)
}

func (array nodePArr) Less(i, j int) bool {
	return array[i].p > array[j].p //从大到小
}

func (array nodePArr) Swap(i, j int) {
	array[i], array[j] = array[j], array[i]
}

func findMaximizedCapital(k, m int, profits, capital []int) int {
	var (
		minCostQ    []*node
		maxProfitsQ []*node
	)
	//所有项目放进加锁数组中，增序排列
	for i := 0; i < len(profits); i++ {
		minCostQ = append(minCostQ, &node{c: capital[i], p: profits[i]})
	}
	sort.Sort(nodeCArr(minCostQ))

	//花费不大于m的项目加入解锁池中
	for i := 0; i < k; i++ {
		for len(minCostQ) > 0 && minCostQ[i].c <= m {
			maxProfitsQ = append(maxProfitsQ, minCostQ[i])
			minCostQ = minCostQ[i+1:]
		}
		if len(maxProfitsQ) == 0 {
			return m
		}
		sort.Sort(nodePArr(maxProfitsQ))
		m += maxProfitsQ[1].p
	}
	return m
}
