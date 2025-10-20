function uniqueOccurrences(arr: number[]): boolean {
  const map = new Map<number, number>();
  for (const a of arr) {
    if (map.has(a)) {
      map.set(a, map.get(a)! + 1);
    } else {
      map.set(a, 0);
    }
  }
  const set = new Set<number>();
  for (const v of map.values()) {
    set.add(v);
  }
  return set.size === map.size;
}
