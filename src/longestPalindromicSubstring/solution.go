//solution for : https://leetcode.com/problems/longest-palindromic-substring/
package longestPalindromicSubstring

func longestPalindrome(s string) string {
	//处理空字符串情况
	if s == "" {
		return ""
	}
	//从最长长度的子串往下找
	for i := len(s); i > 1; i-- {
		//尝试对所有长度的子串判断是否回文字符串
		for j := 0; j < len(s)+1-i; j++ {
			//若是回文字符串，则返回
			if longestString := s[j : j+i]; isPalindrome(longestString) {
				return longestString
			}
		}
	}
	return s[0:1]
}

//判断是否回文字符串
func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

//失败记录：
//case： ""。未考虑空字符串情况

//Runtime: 88 ms, faster than 24.68% of Go online submissions for Longest Palindromic Substring.
//Memory Usage: 2.2 MB, less than 100.00% of Go online submissions for Longest Palindromic Substring.

//总结：时间过长，其实可以从短字符串开始向长字符串判断。（奇偶分开判断）
