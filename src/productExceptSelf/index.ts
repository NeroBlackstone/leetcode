function productExceptSelf(nums: number[]): number[] {
  const len = nums.length;
  // 牢记这种左右分治的思路
  const arrL = new Array<number>(len);
  const arrR = new Array<number>(len);
  arrL[0] = 1;
  arrR[len - 1] = 1;
  for (let i = 0; i < len - 1; i++) {
    arrL[i + 1] = nums[i]! * arrL[i]!;
    const indexR = len - 1 - i;
    arrR[indexR - 1] = nums[indexR]! * arrR[indexR]!;
  }
  for (let i = 0; i < len; i++) {
    arrL[i] = arrL[i]! * arrR[i]!;
  }
  return arrL;
}
// https://leetcode.com/problems/product-of-array-except-self/description/
