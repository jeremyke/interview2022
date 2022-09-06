package main

import "fmt"

//给你一个字符串 s，找到 s 中最长的回文子串。
//babad-->bad cbbd-->bb
//回文：就是左右对称
//思路：（1）从左到右便历以每个字符为中心像两边扩散，找出最长回文（2）从左到右便历2个字符是否相同，相同的以他为中心向两边扩散找出最长回文

func main() {
	s := "abcbb"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	var res = ""
	if s == "" || len(s) == 1 {
		return s
	}
	for i := 0; i < len(s); i++ {
		left, right := i, i
		for left >= 0 && right < len(s) {
			if s[left] != s[right] {
				break
			}
			if right-left+1 > len(res) {
				res = s[left : right+1]
			}
			left--
			right++
		}
	}

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			if len(res) < 2 {
				res = s[i : i+2]
			}
		}
		left, right := i, i+1
		for left >= 0 && right < len(s) {
			if s[left] != s[right] {
				break
			}
			if right-left+1 > len(res) {
				res = s[left : right+1]
			}
			left--
			right++
		}
	}

	return res
}
