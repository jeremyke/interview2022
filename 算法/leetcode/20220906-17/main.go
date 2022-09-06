package main

//编写一个函数来查找字符串数组中的最长公共前缀。
//
//如果不存在公共前缀，返回空字符串 ""。

func main() {
	str := "23"
	println(letterCombinations(str))
}

func letterCombinations(digits string) []string {
	numStringMap := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	i := 0
	if len(digits) > i {

		for k, v := range numStringMap[string(digits[i])] {
			i--
			if len(digits) > i {

			}
		}
	}

}
