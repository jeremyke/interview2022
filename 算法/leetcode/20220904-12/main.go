package main

import "fmt"

func main() {
	num := 1994
	fmt.Println(intToRoman(num))
}

//整数转罗马数字
func intToRoman(num int) string {
	var iRMapArr = []struct {
		numKey int
		value  string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	resStr := ""
	for _, v := range iRMapArr {
		for num >= v.numKey {
			resStr = resStr + v.value
			num = num - v.numKey
		}
		if num == 0 {
			break
		}
	}
	return resStr
}
