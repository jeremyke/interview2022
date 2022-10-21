package main

func process(str []byte, i int) int { //str待转化的数字字符串，i之前位置的元素已经转化好了，i之后的元素一共多少种转化方法
	if i == len(str) {
		return 1
	}
	//str[i]=0
	if str[i] == '0' {
		return 0
	}

	//str[i]=1
	if str[i] == '1' {
		res := process(str, i+1) //i单独转化，返回i后面元素转化结果种数
		if i+1 < len(str) {
			res += process(str, i+2) //i和i+1结合转化，返回i+1后面的元素的转化结果种数
		}
		return res
	}

	//str[i]=2
	if str[i] == '2' {
		res := process(str, i+1)                                  //i单独转化，返回i后面元素转化结果种数
		if i+1 < len(str) && str[i+1] >= '0' && str[i+1] <= '6' { //i和i+1结合转化不超过26，后续几种方法
			res += process(str, i+2)
		}
		return res
	}

	//str[i]介于3-9之间，只有一种决定
	return process(str, i+1)
}
