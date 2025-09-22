function moveZeroes(nums: number[]): void {
  let zeroPointer = nums.indexOf(0);
  if (zeroPointer !== -1) {
    for (let i = 0; i < nums.length; i++) {
      if (nums[i] !== 0 && i > zeroPointer) {
        nums[zeroPointer] = nums[i]!;
        zeroPointer++;
      }
    }
    while (zeroPointer < nums.length) {
      nums[zeroPointer] = 0;
      zeroPointer++;
    }
  }
}
