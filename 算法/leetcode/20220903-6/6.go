package main

func main() {

}

func convert(s string, numRows int) string {
	sLen := len(s)
	if numRows == 1 || numRows >= sLen {
		return s
	}
	t := numRows*2 - 2 //每个循环周期的字符个数
	// 每个周期占用numRows-2+1列
	c := (sLen + t - 1) / t * (numRows - 1) //二维数组的列数
	//构建二维数组
	mat := make([][]byte, numRows)
	for i := range mat {
		mat[i] = make([]byte, c)
	}

	x, y := 0, 0
	for i, ch := range s {
		mat[x][y] = byte(ch)
		//索引求周期字符个数的余大于行数-1就下移，否则上移动
		if i%t < numRows-1 {
			x++ // 向下移动
		} else {
			x--
			y++ // 向右上移动
		}
	}
	//二维数组去空转为字符串
	ans := make([]byte, 0, sLen)
	for _, row := range mat {
		for _, ch := range row {
			if ch > 0 {
				ans = append(ans, ch)
			}
		}
	}
	return string(ans)
}
