package medianofTwoSortedArrays

class Solution {
    fun findMedianSortedArrays(nums1: IntArray, nums2: IntArray): Double {
        val nums=nums1+nums2
        nums.sort()
        val size=nums.size
        val halfSize=size/2
        return if (size%2!=0)
            nums[halfSize].toDouble()
        else
            (nums[halfSize]+nums[halfSize-1])/2.0
    }
}

fun main() {
    val nums1= intArrayOf(1,3)
    val nums2= intArrayOf(2)
    val solution=Solution()
    print(solution.findMedianSortedArrays(nums1,nums2))
}