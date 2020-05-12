package stringToInteger

import (
	"math"
	"strconv"
	"strings"
	"unicode"
)

func myAtoi(str string) int {
	//去除字符串的空格
	str = strings.TrimSpace(str)
	//判空，为空直接返回0
	if len(str) == 0 {
		return 0
	}
	//第一个字符就非法
	if invalidRune([]rune(str)[0]) {
		return 0
	}
	//从一开始找到第一个不为数字的字符索引
	index := -1
	for i := 1; i < len(str); i++ {
		if !isANumber(rune(str[i])) {
			index = i
			break
		}
	}
	//把合法部分截取下来
	if index != -1 {
		str = str[0:index]
	}
	//字符串转为整型数
	integer, _ := strconv.Atoi(str)
	//整型数超过范围，取最大范围。
	if integer > math.MaxInt32 {
		return math.MaxInt32
	} else if integer < math.MinInt32 {
		return math.MinInt32
	}
	return integer
}

//字符串左边只能为数字，+-号
func invalidRune(r rune) bool {
	return !isANumber(r) && r != '-' && r != '+'
}

//判断字符是否是一个数字
func isANumber(r rune) bool {
	return unicode.IsNumber(r)
}

//Runtime: 4 ms, faster than 46.91% of Go online submissions for String to Integer (atoi).
//Memory Usage: 2.3 MB, less than 25.00% of Go online submissions for String to Integer (atoi).

//总结：这道题边界情况特别多，建议老老实实截取字符串做，思路是找到第一个非法到字符索引，截取即可。
