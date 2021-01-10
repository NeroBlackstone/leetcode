package searchInsertPosition

import "sort"

// https://leetcode.com/problems/search-insert-position/
func searchInsert(nums []int, target int) int {
	return sort.SearchInts(nums, target)
}

// Runtime: 4 ms, faster than 86.17% of Go online submissions for Search Insert Position.
//Memory Usage: 3.1 MB, less than 100.00% of Go online submissions for Search Insert Position.
//总结：直接使用标准库作弊
