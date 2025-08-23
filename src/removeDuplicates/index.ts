// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/
function removeDuplicates(nums: number[]): number {
  // fast代表待比较元素
  // slow代表待写入位置
  let slowPointer = 0;
  for (let fastPointer = 0; fastPointer < nums.length; fastPointer++) {
    // fast和slow-2不相等，说明可以写入slow
    if (slowPointer < 2 || nums[slowPointer - 2] !== nums[fastPointer]) {
      nums[slowPointer] = nums[fastPointer]!;
      slowPointer++;
    }
  }
  return slowPointer;
}

// 知识点，双指针。
