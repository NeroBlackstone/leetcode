function majorityElement(nums: number[]): number {
  // 总共出现n次
  let count = 0;
  // 候选数字
  let candidate = 0;
  for (const num of nums) {
    if (count === 0) {
      candidate = num;
    }
    count += candidate === num ? 1 : -1;
  }
  return candidate;
}

// 知识点Boyer–Moore Majority Vote Algorithm（摩尔投票算法）
// https://leetcode.com/problems/majority-element/description/
