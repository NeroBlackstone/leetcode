function findMaxAverage(nums: number[], k: number): number {
  let max = nums.slice(0, k).reduce((a, b) => a + b, 0);
  let current = max;
  for (let i = 1; i <= nums.length - k; i++) {
    current = current - nums[i - 1]! + nums[i + k - 1]!;
    max = Math.max(current, max);
  }
  return max / k;
}
