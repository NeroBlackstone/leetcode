package lengthOfLastWord

import "strings"

//https://leetcode.com/problems/length-of-last-word/
func lengthOfLastWord(s string) int {
	if s == " " {
		return 0
	}
	s = strings.TrimSpace(s)
	length := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			length++
		} else {
			break
		}
	}
	return length
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Length of Last Word.
// Memory Usage: 2.1 MB, less than 48.33% of Go online submissions for Length of Last Word.
// 总结：看到题解里有人用strings.Split函数，虽然很方便但我感觉效率不会很高，因为要处理很多无用数据。
