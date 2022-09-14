package main

import (
	"fmt"
	"sort"
)

//给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
//
//字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。
//

//思路：

func main() {
	strArr := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strArr))
}

func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}
	//便历字符串数组，把字母异位词字符串存入同一个map里面
	for _, str := range strs {
		s := []byte(str)
		//每个字符串的字符进行排序，互为字母异位词字符串，具有相同的字符，排序后的字符串相同
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedStr := string(s)
		mp[sortedStr] = append(mp[sortedStr], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}
