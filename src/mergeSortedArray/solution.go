package mergeSortedArray

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Merge Sorted Array.
//Memory Usage: 2.3 MB, less than 100.00% of Go online submissions for Merge Sorted Array.
