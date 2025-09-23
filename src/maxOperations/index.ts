function maxOperations(nums: number[], k: number): number {
  nums.sort((a, b) => a - b);
  let left = 0;
  let right = nums.length - 1;
  let count = 0;
  while (left < right) {
    const leftN = nums[left]!;
    const rightN = nums[right]!;
    if (leftN + rightN === k) {
      count++;
      left++;
      right--;
    } else if (leftN + rightN < k) {
      left++;
    } else {
      right--;
    }
  }
  return count;
}
