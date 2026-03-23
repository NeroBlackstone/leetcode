//示例 1：
//输入：nums = [1,1,1], k = 2
//输出：2
function subarraySum(nums: number[], k: number): number {
  let result = 0;
  const numLen = nums.length;
  for (let i = 0; i < numLen; i++) {
    let sum = 0;
    for (let j = i; j < numLen; j++) {
      sum += nums[j]!;
      if (sum === k) {
        result++;
      }
    }
  }
  return result;
}
