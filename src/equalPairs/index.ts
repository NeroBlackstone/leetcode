function equalPairs(grid: number[][]): number {
  const map = new Map<string, number>();

  const rows = grid.map((g) => g.join());

  for (const r of rows) {
    map.set(r, (map.get(r) ?? 0) + 1);
  }
  let count = 0;
  for (let i = 0; i < grid.length; i++) {
    const column = grid.map((g) => g[i]).join();
    count = count + (map.get(column) ?? 0);
  }

  return count;
}
