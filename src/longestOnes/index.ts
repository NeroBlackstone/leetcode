function longestOnes(nums: number[], k: number): number {
  let zeroCount = 0;
  let maxLength = 0;
  let l = 0;
  for (let r = 0; r < nums.length; r++) {
    if (nums[r] === 0) {
      zeroCount++;
    }
    while (zeroCount > k && l <= r) {
      if (nums[l] === 0) {
        zeroCount--;
      }
      l++;
    }
    maxLength = Math.max(maxLength, r - l + 1);
  }
  return maxLength;
}
