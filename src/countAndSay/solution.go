package countAndSay

import "strconv"

// https://leetcode.com/problems/count-and-say/
func countAndSay(n int) string {
	s := "1"
	for i := 1; i < n; i++ {
		s = nextSay(s)
	}
	return s
}

// Get next say
func nextSay(s string) string {
	nextStr := ""
	nowNum := '0'
	nowCount := 0
	for _, e := range s {
		if nowNum != e {
			if nowNum != '0' {
				nextStr += getSubStr(nowCount, nowNum)
			}
			nowNum = e
			nowCount = 1
		} else {
			nowCount++
		}
	}
	nextStr += getSubStr(nowCount, nowNum)
	return nextStr
}

// Get ∣count,digit∣ sub-string
func getSubStr(count int, num int32) string {
	return strconv.Itoa(count) + string(num)
}

// Runtime: 12 ms, faster than 23.08% of Go online submissions for Count and Say.
//Memory Usage: 7.1 MB, less than 10.44% of Go online submissions for Count and Say.
// 总结：这是暴力求解，比较慢，但是易于理解。也可以使用recursion去做。没有太多本质差别。
