function findDifference(nums1: number[], nums2: number[]): number[][] {
  const set1 = new Set<number>();
  const set2 = new Set<number>();
  for (const n of nums1) {
    set1.add(n);
  }
  for (const n of nums2) {
    set2.add(n);
  }
  const diff1 = [];
  const diff2 = [];
  for (const n of set1) {
    if (!set2.has(n)) {
      diff1.push(n);
    }
  }
  for (const n of set2) {
    if (!set1.has(n)) {
      diff2.push(n);
    }
  }
  return [diff1, diff2];
}
