function largestAltitude(gain: number[]): number {
  let maxCount = 0;
  let sum = 0;
  for (const g of gain) {
    sum += g;
    maxCount = Math.max(sum, maxCount);
  }
  return maxCount;
}
