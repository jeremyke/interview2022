package main

import "math"

func num1(n int) int {
	if n < 1 {
		return 1
	}
	record := make([]int, n) //record[i]-->第i行的皇后放在了第几列
	return process1(0, n, record)
}

func process1(i, n int, record []int) int { //i表示行数；n一共多少行；返回合法的摆法
	if i == n {
		return 1
	}
	res := 0
	for j := 0; j < n; j++ { //在i行尝试所有列
		//当前i行的皇后，放在j列，会不会和之前0--i-1的皇后，共行、共列或者共斜线，是则为无效；否则为有效
		if isValid(record, i, j) {
			record[i] = j
			res += process1(i+1, n, record)
		}
	}
	return res
}

//检查是否合法（共行、共列或者共斜线）
func isValid(record []int, i, j int) bool { //i表示行数；j表示在i行尝试的列数
	for k := 0; k < i; k++ { //之前的某个k行的皇后
		if j == record[k] || math.Abs(float64(record[k]-j)) == math.Abs(float64(i-k)) {
			return false
		}
	}
	return true
}
