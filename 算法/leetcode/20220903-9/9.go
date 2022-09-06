package main

import (
	"fmt"
	"strconv"
)

//给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

func main() {
	sss := 10

	fmt.Println(isPalindrome(sss))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	xStr := strconv.Itoa(x)
	flag := true
	var left, right int
	if len(xStr)%2 == 0 {
		left, right = len(xStr)/2-1, len(xStr)/2
		if xStr[left] != xStr[right] {
			flag = false
			return flag
		}

	} else {
		left, right = (len(xStr)-1)/2, (len(xStr)-1)/2
	}
	for left >= 0 && right < len(xStr) {
		if xStr[left] != xStr[right] {
			flag = false
			break
		}
		left--
		right++
	}
	return flag
}
