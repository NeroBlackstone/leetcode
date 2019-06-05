package twoSum

class Solution {
    fun twoSum(nums: IntArray, target: Int): IntArray {
        nums.sliceArray(0 until  nums.size-1).forEachIndexed{ indexFirst,first->
            nums.sliceArray(indexFirst+1 until nums.size).forEachIndexed{indexSecond,second->
                if (first+second==target)
                    return intArrayOf(indexFirst,indexSecond+indexFirst+1)
            }
        }
        return intArrayOf(0,0)
    }
}

fun main() {
    val solution=Solution()
    val nums=intArrayOf(2,7,11,15)
    val target=9
    val out=solution.twoSum(nums,target)
    out.forEach { print("$it ") }
}