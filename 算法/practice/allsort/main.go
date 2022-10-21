package main

import "fmt"

func childStr(str string) {
	strArr := []byte(str)
	process(strArr, 0, []byte{}, map[byte]int{})
}

func process(strArr []byte, i int, res []byte, resMap map[byte]int) { //第i个字符选哪个字符；res之前的选择形成的数组；resMap之前的选择形成的map
	if i == len(strArr) {
		fmt.Println(string(res))
		return
	}
	for j := 0; j < len(strArr); j++ {
		if _, ok := resMap[strArr[j]]; !ok {
			res = append(res, strArr[j])
			resMapTmp := cloneMap(resMap) //复制map,保证每一步程序这里只记录已经选择的字符集合
			resMapTmp[strArr[j]] = 1
			process(strArr, i+1, res, resMapTmp)
			res = res[0 : len(res)-1] //上一步执行完后，去掉上一步选择的字符，重新选择字符
		}
	}
}

func cloneMap(tags map[byte]int) map[byte]int {
	cloneTags := make(map[byte]int)
	for k, v := range tags {
		cloneTags[k] = v
	}
	return cloneTags
}

func main() {
	childStr("ab")
}
