package longestCommonPrefix

import (
	"unicode"
)

// https://leetcode.com/problems/longest-common-prefix/
func longestCommonPrefix(strs []string) string {
	// 0 <= strs.length <= 200
	strsLength := len(strs)
	if strsLength > 200 || strsLength <= 0 {
		return ""
	}

	loopTimes := 200
	for _, str := range strs {
		// 0 <= strs[i].length <= 200
		strLength := len(str)
		if strLength > 200 || strLength <= 0 {
			return ""
		}
		if strLength < loopTimes {
			loopTimes = strLength
		}
		// strs[i] consists of only lower-case English letters.
		for _, v := range str {
			if !unicode.IsLower(v) {
				return ""
			}
		}
	}

	var prefix string
out:
	for flag, i := true, 0; flag && i < loopTimes; i++ {
		thisRune := strs[0][i]
		for _, str := range strs {
			if thisRune != str[i] {
				flag = false
				break out
			}
		}
		prefix += string(thisRune)
	}
	return prefix
}

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Longest Common Prefix.
//Memory Usage: 2.4 MB, less than 23.05% of Go online submissions for Longest Common Prefix.
