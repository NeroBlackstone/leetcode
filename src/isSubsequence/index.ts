function isSubsequence(s: string, t: string): boolean {
  let sPointer = 0;
  let tPointer = 0;
  while (sPointer < s.length && tPointer < t.length) {
    if (t[tPointer] === s[sPointer]) {
      sPointer++;
    }
    tPointer++;
  }
  return sPointer === s.length;
}
