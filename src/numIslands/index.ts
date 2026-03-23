function numIslands(grid: string[][]): number {
  let count = 0;
  const rows = grid.length;
  const cols = grid[0]!.length;

  //发现大陆，淹没陆地
  function dfs(r: number, c: number) {
    // 在边界外，直接返回
    if (r < 0 || r >= rows || c < 0 || c >= cols) {
      return;
    }
    // 在边界内，看是不是陆地，如果是水就返回，是陆地就淹没
    if (grid[r]![c] === "0") {
      return;
    } else {
      grid[r]![c] = "0";
    }

    // 向外搜索
    dfs(r + 1, c);
    dfs(r - 1, c);
    dfs(r, c + 1);
    dfs(r, c - 1);
  }

  for (let i = 0; i < rows; i++) {
    for (let j = 0; j < cols; j++) {
      if (grid[i]![j] === "1") {
        count++;
        dfs(i, j);
      }
    }
  }
  return count;
}
