package main

//编写一个函数来查找字符串数组中的最长公共前缀。
//
//如果不存在公共前缀，返回空字符串 ""。

func main() {
	//str := []string{"flower", "flow", "flight"}
	str := []string{"a"}
	//str := []string{"abab", "aba", ""}
	//str := []string{"ac", "ac", "a", "a"}
	//str := []string{"ac", "a"}
	//str := []string{"froolic", "froolic", "fraodf", "frloadsf", "frlozcv"}
	println(longestCommonPrefix(str))
}

//最长公共前缀
func longestCommonPrefix(strs []string) string {
	comPre := ""
	if len(strs) == 0 {
		return comPre
	}
	if len(strs) == 1 {
		return strs[0]
	}
	//找出前2个字符串的公共前缀
	for i := 0; i < len(strs[0]) && i < len(strs[1]); i++ {
		if strs[0][i] != strs[1][i] {
			break
		}
		comPre = comPre + string(strs[0][i])
	}
	//var tmpComPre string
	for j := 2; j < len(strs); j++ {
		if len(strs[j]) < len(comPre) {
			comPre = comPre[:len(strs[j])]
		}
		if len(strs[j]) >= len(comPre) {
			for k := 0; k < len(comPre); k++ {
				if comPre[k] != strs[j][k] {
					comPre = comPre[:k]
					break
				} else {
					continue
				}
			}
		}
	}

	return comPre
}
