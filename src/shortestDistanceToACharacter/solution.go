package shortestDistanceToACharacter

import (
	"math"
	"github.com/NeroBlackstone/leetcode/collection"
)

// https://leetcode.com/problems/shortest-distance-to-a-character/
func shortestToChar(s string, c byte) []int {
	l := len(s)
	ans := make([]int,l)
	index := math.MinInt32/2
	for i,e := range s{
		if byte(e) == c {
			index=i
		}
		// for rune before C in String S, the distance will very BIG
		ans[i]=i-index
	}
	index=math.MaxInt32/2
	for i:=l-1;i>=0;i--{
		if s[i]==c{
			index=i
		}
		ans[i]=collection.Min(ans[i],index-i)
	}
	return ans
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Shortest Distance to a Character.
// Memory Usage: 2.4 MB, less than 100.00% of Go online submissions for Shortest Distance to a Character.