package main

import (
	"fmt"
)

//给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

func main() {
	s := "ccb"
	//fmt.Printf("res=%v", isReaped(s))
	fmt.Println(lengthOfLongestSubstring(s))
}

//滑动窗口解法
func lengthOfLongestSubstring(s string) int {
	var res int
	//空字符串，最大无重复字符串长度为0
	if s == "" {
		return res
	}
	if len(s) == 1 {
		res = 1
		return res
	}
	//非空时
	window := map[byte]bool{
		s[0]: true,
	}
	var l = 0
	for r := 1; r < len(s); {
		if !window[s[r]] {
			if r-l+1 > res {
				res = r - l + 1
			}
			window[s[r]] = true
			r++
		} else {
			delete(window, s[l])
			l++
		}

	}
	return res
}

//暴力解法
//func lengthOfLongestSubstring(s string) int {
//	if s == "" {
//		return 0
//	}
//	var (
//		res   int
//		aChan = make(chan int, len(s)*len(s))
//	)
//	for m := 2; m <= len(s); m++ {
//		for n := 0; n <= len(s)-m; n++ {
//			go func(n, m int) {
//				if !isReaped(s[n : n+m]) {
//					aChan <- m
//					if m == len(s) && n == len(s)-m {
//						close(aChan)
//					}
//				}
//			}(n, m)
//
//		}
//	}
//	time.Sleep(3)
//	lenth := len(aChan)
//	for i := 0; i < lenth; i++ {
//		num, ok := <-aChan
//		if !ok {
//			break
//		}
//		if num > res {
//			res = num
//		}
//		//fmt.Printf("i=%v--num=%v\n", i, num)
//
//	}
//	if res == 0 {
//		return 1
//	} else {
//		return res
//	}
//}

//func isReaped(str string) bool {
//	for i := 0; i < len(str)-1; i++ {
//		for j := i + 1; j < len(str); j++ {
//			if str[i] == str[j] {
//				return true
//			}
//		}
//	}
//	return false
//}
