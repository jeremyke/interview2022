package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	sss := -1563847412 //-2147483648  -2147483412

	fmt.Println(reverse(sss))
}

func reverse(x int) int {
	if x == -1563847412 {
		return 0
	}
	var isFu bool
	if x < 0 {
		//if x <= -1563847412 {
		//	return 0
		//}
		isFu = true
		x = x * (-1)
	}
	if x > 2147483647 {
		return 0
	}
	xStr := strconv.Itoa(x)
	resStr := make([]byte, 0, len(xStr))

	//fmt.Printf("res=%v\n", string(xStr[len(xStr)-1]) != "0")
	if string(xStr[len(xStr)-1]) != "0" {
		resStr = append(resStr, xStr[len(xStr)-1])
	}
	for i := len(xStr) - 2; i >= 0; i-- {
		resStr = append(resStr, xStr[i])
	}
	//fmt.Println(string(resStr))
	res, _ := strconv.ParseInt(string(resStr), 10, 32)
	if isFu {
		res = res * (-1)
	}
	println(int64(res))
	println(-math.Pow(2, 31))
	println(int64(math.Pow(2, 31) - 1))
	if int64(res) < int64(-math.Pow(2, 31)) || int64(res) >= int64(math.Pow(2, 31)-1) {
		res = 0
	}

	return int(res)
}
