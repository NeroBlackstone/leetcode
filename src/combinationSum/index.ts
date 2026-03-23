// 输入：candidates = [2,3,6,7], target = 7
// 输出：[[2,2,3],[7]]
function combinationSum(candidates: number[], target: number): number[][] {
  const result: number[][] = [];
  const path: number[] = [];
  function backtrack(start: number, sum: number) {
    if (sum === target) {
      result.push([...path]);
      return;
    }
    if (sum > target) {
      return;
    }
    for (let i = start; i < candidates.length; i++) {
      const c = candidates[i]!;
      path.push(c);
      backtrack(i, sum + c);
      path.pop();
    }
  }
  backtrack(0, 0);
  return result;
}
