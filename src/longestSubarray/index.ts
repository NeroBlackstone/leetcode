function longestSubarray(nums: number[]): number {
  let maxLength = 0;
  let l = 0;
  let zeroCount = 0;
  for (let r = 0; r < nums.length; r++) {
    if (nums[r] === 0) {
      zeroCount++;
    }
    while (zeroCount > 1 && l <= r) {
      if (nums[l] === 0) {
        zeroCount--;
      }
      l++;
    }
    maxLength = Math.max(maxLength, r - l + 1);
  }
  return maxLength - 1;
}
