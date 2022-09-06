package main

//func main() {
//	nums := []int{0, 1, 3, 6, 9, 2, 7}
//	target := 9
//	res := getAdd(nums, target)
//	fmt.Println(res)
//}

func getAdd(nums []int, target int) [][]int {
	nums_count := len(nums)
	if nums_count <= 0 {
		return [][]int{}
	}
	var res [][]int
	for i := 0; i < nums_count; i++ {
		for j := i + 1; j < nums_count; j++ {
			if nums[i]+nums[j] == target {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}
