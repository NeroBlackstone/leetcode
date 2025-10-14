function pivotIndex(nums: number[]): number {
  const sum = nums.reduce((a, b) => a + b, 0);
  let sumL = 0;
  for (let i = 0; i < nums.length; i++) {
    if (sumL * 2 === sum - nums[i]!) {
      return i;
    }
    sumL += nums[i]!;
  }
  return -1;
}
