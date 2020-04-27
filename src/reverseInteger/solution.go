package reverseInteger

import (
	"math"
	"strconv"
)

func reverse(x int) int {
	//处理输入的边界情况
	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}

	//算出x的绝对值
	abs := abs(x)
	//绝对值转字符串
	s := strconv.Itoa(abs)
	result, _ := strconv.Atoi(reverseStr(s))
	if x < 0 {
		result = -result
	}

	//处理输出的边界情况
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	return result
}

func reverseStr(s string) string {
	bytes := []byte(s)
	for i := 0; i < len(bytes)/2; i++ {
		//要交换的index
		swapI := len(bytes) - 1 - i
		//交换
		c := bytes[i]
		bytes[i] = bytes[swapI]
		bytes[swapI] = c
	}
	return string(bytes)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Integer.
//Memory Usage: 2.2 MB, less than 33.33% of Go online submissions for Reverse Integer.

//总结，这道题要注意处理边界情况，不仅要考虑输入的边界，也要考虑输出的边界。
