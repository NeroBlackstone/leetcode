function jump(nums: number[]): number {
  let maxReach = 0;
  let count = 0;
  let end = 0;
  // 注意这里是<nums.length -1,因为一定可以到最后一个元素，所以你不用考虑最后一个元素怎么跳
  for (let i = 0; i < nums.length - 1; i++) {
    maxReach = Math.max(maxReach, i + nums[i]!);
    if (i === end) {
      count++;
      end = maxReach;
    }
  }
  return count;
}

// https://leetcode.com/problems/jump-game-ii/
