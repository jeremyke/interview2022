package main

import "fmt"

func hanoi(n int) {
	if n > 0 {
		process(n, "左", "右", "中")
	}
}

//1~i圆盘从左移动到右边
func process(i int, start, end, other string) {
	if i == 1 {
		fmt.Printf("move 1 from %s to %s\n", start, end)
	} else {
		process(i-1, start, other, end)                      //把1~i-1号元素，移动到other
		fmt.Printf("move %v from %s to %s\n", i, start, end) //把i号元素移动到end
		process(i-1, other, end, start)                      //再把把1~i-1号元素，移动到end
	}
}

func main() {
	hanoi(3)
}
