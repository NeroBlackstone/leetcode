// 输入：nums = [1,2,3]
// 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
function subsets(nums: number[]): number[][] {
  const result: number[][] = [[]];
  for (const n of nums) {
    const size = result.length;
    for (let i = 0; i < size; i++) {
      result.push([...result[i]!, n]);
    }
  }
  return result;
}
