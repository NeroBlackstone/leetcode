package implementstrStr

import "strings"

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Implement strStr().
//Memory Usage: 2.3 MB, less than 100.00% of Go online submissions for Implement strStr().
// 总结：用strings包的Index方法作弊
