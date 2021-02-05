package longestHarmoniousSubsequence

import "github.com/NeroBlackstone/leetcode/collection"

// https://leetcode.com/problems/longest-harmonious-subsequence/
func findLHS(nums []int) int {
	m,res:=make(map[int]int),0
	for _,e:=range nums{
		m[e]=m[e]+1
		if m[e+1]!=0{
			res=collection.Max(res,m[e]+m[e+1])
		}
		if m[e-1]!=0{
			res=collection.Max(res,m[e]+m[e-1])
		}
	}
	return res
}

//Runtime: 64 ms, faster than 65.38% of Go online submissions for Longest Harmonious Subsequence.
//Memory Usage: 6.6 MB, less than 78.85% of Go online submissions for Longest Harmonious Subsequence.
//Summary:Single Loop and HashMap