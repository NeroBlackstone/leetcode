function permute(nums: number[]): number[][] {
  const res: number[][] = [];
  const path: number[] = [];

  // 记录已经使用过的元素
  const used = new Set<number>();

  function backtrack() {
    if (path.length === nums.length) {
      res.push([...path]);
      return;
    }

    for (const n of nums) {
      // 如果已经使用过
      if (used.has(n)) continue;

      // 做选择
      path.push(n);
      used.add(n);

      backtrack();

      // 回溯
      path.pop();
      used.delete(n);
    }
  }

  backtrack();
  return res;
}
