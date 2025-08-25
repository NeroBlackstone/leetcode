// https://leetcode.com/problems/rotate-array/
function rotate(nums: number[], k: number): void {
  const numsLength = nums.length;
  const endIndex = numsLength - 1;
  // 注意这里k可能大于nums.length
  const shift = k % numsLength;
  reverseNumArray(nums, 0, endIndex);
  reverseNumArray(nums, 0, shift - 1);
  reverseNumArray(nums, shift, endIndex);
}

function reverseNumArray(nums: number[], start: number, end: number) {
  // 注意这里要考虑start != 0的情况,不能直接start++<end/2来判断
  // 注意这里要start < end，不能直接判断start === end不然会出现start=5,end=6的情况
  while (start < end) {
    // 注意这里可以通过解构赋值少写一个temp变量
    [nums[start], nums[end]] = [nums[end]!, nums[start]!];
    start++;
    end--;
  }
}
