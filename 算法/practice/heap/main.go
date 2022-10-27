package main

import "fmt"

func main() {
	var arr = []int{1, 3, 7, 6, 4, 2, 8}
	//newArr := []int{1}
	//for i := 1; i < len(arr); i++ {
	//	newArr = append(newArr, arr[i])
	//	heapInsert(newArr, len(newArr)-1)
	//}
	//fmt.Println(newArr)

	//heapify(arr, 0, 7)
	//fmt.Println(arr)
	//edit(newArr, 2, 9)
	heapSort(arr)
	fmt.Println(arr)
}

//
// heapInsert
// @Description: 添加节点
// @Author: maxwell.ke
// @time 2022-10-27 15:49:49
// @param arr
// @param index
//
func heapInsert(arr []int, index int) {
	for arr[index] > arr[(index-1)/2] { //比父节点大，就和父节点交换
		swap(arr, index, (index-1)/2)
		index = (index - 1) / 2
	}
}

//
// heapify
// @Description: 删除节点
// @Author: maxwell.ke
// @time 2022-10-27 15:50:04
// @param arr
// @param index
// @param heapSize
//
func heapify(arr []int, index, heapSize int) {
	left := index*2 + 1
	for left < heapSize {
		//取出2个子节点中最大的节点位置
		biggerChild := left
		if left+1 < heapSize && arr[left+1] > arr[left] {
			biggerChild = left + 1
		}
		//把当前节点的数和子节点较大的数比较
		if arr[biggerChild] < arr[index] {
			biggerChild = index
		}
		if biggerChild == index {
			break
		}
		//交换位置
		swap(arr, biggerChild, index)
		index = biggerChild
		left = index*2 + 1
	}
}

//
// edit
// @Description: 修改堆的某个数，依然保存堆结构
// @Author: maxwell.ke
// @time 2022-10-27 16:05:00
// @param arr
// @param index
// @param value
//
func edit(arr []int, index, value int) {
	if index >= len(arr) || index <= 0 {
		return
	}
	heapSize := len(arr)
	if value < arr[index] {
		arr[index] = value
		heapify(arr, index, heapSize)
		return
	}
	if value > arr[index] {
		arr[index] = value
		heapInsert(arr, index)
		return
	}
	return
}

//
// heapSort
// @Description: 堆排序
// @Author: maxwell.ke
// @time 2022-10-27 16:05:25
// @param arr
//
func heapSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	//把数组放入一个大根堆里面
	for i := 0; i < len(arr); i++ {
		heapInsert(arr, i)
	}
	heapSize := len(arr)
	//0位置的数和最后位置数交换
	swap(arr, 0, heapSize-1)
	heapSize--
	//再把0位置数heapSize
	for heapSize > 0 {
		heapify(arr, 0, heapSize)
		swap(arr, 0, heapSize-1)
		heapSize--
	}
}

func swap(arr []int, a, b int) {
	arr[a] = arr[a] ^ arr[b]
	arr[b] = arr[a] ^ arr[b]
	arr[a] = arr[a] ^ arr[b]
}
