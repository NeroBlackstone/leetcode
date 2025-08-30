function canJump(nums: number[]): boolean {
  let maxReachIndex = 0;
  for (let i = 0; i <= maxReachIndex; i++) {
    const canReachIndex = nums[i]! + i;
    if (canReachIndex >= nums.length - 1) {
      return true;
    }
    if (canReachIndex > maxReachIndex) {
      maxReachIndex = canReachIndex;
    }
  }
  return false;
}
