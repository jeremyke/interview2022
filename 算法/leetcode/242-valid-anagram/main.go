package main

import "fmt"

//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//
//注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。
//

func main() {
	//s, t := "anagram", "nagaram"
	s, t := "a", "ab"
	fmt.Println(isAnagram(s, t))
}

func isAnagram(s string, t string) bool {
	sMap, sLens := strMap([]byte(s))
	tMap, tLens := strMap([]byte(t))
	flag := true
	if sLens > tLens {
		for k, v := range sMap {
			if tMap[k] != v {
				flag = false
				break
			}
		}
	} else {
		for k, v := range tMap {
			if sMap[k] != v {
				flag = false
				break
			}
		}
	}
	return flag
}

func strMap(str []byte) (aMap map[byte]int, lens int) {
	lens = len(str)
	aMap = make(map[byte]int, lens)
	for _, v := range str {
		aMap[v]++
	}
	return aMap, lens
}
