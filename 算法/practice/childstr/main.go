package main

import "fmt"

func childStr(str string) {
	strArr := []byte(str)
	process(strArr, 0, []byte{})
}

func process(strArr []byte, i int, res []byte) { //第i个字符要或者不要；res之前的选择形成的数组
	if i == len(strArr) {
		fmt.Println(string(res))
		return
	}
	//加第i个字符的情况
	resKeep := res
	resKeep = append(resKeep, strArr[i])
	process(strArr, i+1, resKeep)
	//不加第i个字符的情况
	resNoInclude := res
	process(strArr, i+1, resNoInclude)
}

func main() {
	childStr("abc")
}
