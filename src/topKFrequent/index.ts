function topKFrequent(nums: number[], k: number): number[] {
  const map = new Map<number, number>();
  for (const n of nums) {
    map.set(n, (map.get(n) ?? 0) + 1);
  }
  const bucket: number[][] = Array(nums.length + 1)
    .fill(null)
    .map(() => []);
  for (const [key, value] of map) {
    bucket[value]?.push(key);
  }
  const res = [];
  for (let i = nums.length; i > 0; i--) {
    for (const j of bucket[i]!) {
      res.push(j);
      if (res.length >= k) {
        return res;
      }
    }
  }
  return res;
}
//  https://leetcode.com/problems/top-k-frequent-elements/
