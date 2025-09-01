function hIndex(citations: number[]): number {
  let h = 0;
  for (const c of citations.sort((a, b) => b - a)) {
    if (c > h) {
      h++;
    } else {
      break;
    }
  }
  return h;
}
// https://leetcode.com/problems/h-index/description/
