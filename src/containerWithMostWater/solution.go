package containerWithMostWater

import "github.com/NeroBlackstone/leetcode/collection"

//https://leetcode.com/problems/container-with-most-water/
func maxArea(height []int) int {
	max, l, r := 0, 0, len(height)-1
	for l < r {
		lh, rh := height[l], height[r]
		max = collection.Max(max, collection.Min(lh, rh)*(r-l))
		if lh < rh {
			l++
		} else {
			r--
		}
	}
	return max
}

// Runtime: 24 ms, faster than 36.44% of Go online submissions for Container With Most Water.
// Memory Usage: 6.1 MB, less than 5.21% of Go online submissions for Container With Most Water.
// 总结：双指针思想
