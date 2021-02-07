package findFirstAndLastPositionOfElementInSortedArray

import "sort"

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
func searchRange(nums []int, target int) []int {
	l:=sort.SearchInts(nums,target)
	if l==len(nums)||nums[l]!=target{
		return []int{-1,-1}
	}
	r:=sort.SearchInts(nums,target+1)-1
	return []int{l,r}
}