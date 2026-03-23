function longestConsecutive(nums: number[]): number {
  if (nums.length === 0) return 0;
  nums.sort((a, b) => a - b);
  let max = 1;
  let longestStrike = 1;
  let preNum = nums[0]!;
  for (const n of nums) {
    if (n === preNum + 1) {
      longestStrike++;
    } else if (n === preNum) {
    } else {
      max = Math.max(longestStrike, max);
      longestStrike = 1;
    }
    preNum = n;
  }
  max = Math.max(longestStrike, max);
  return max;
}
