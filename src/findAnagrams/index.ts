function findAnagrams(s: string, p: string): number[] {
  const pLen = p.length;
  if (s.length < pLen) {
    return [];
  }
  const pFrequent = getStringFrequentArray(p);
  const sFrequent = getStringFrequentArray(s.substring(0, pLen));
  const result = isEqualArray(pFrequent, sFrequent) ? [0] : [];
  for (let i = 1; i <= s.length - pLen; i++) {
    sFrequent[charToNum(s[i - 1]!)]!--;
    sFrequent[charToNum(s[i + pLen - 1]!)]!++;
    if (isEqualArray(pFrequent, sFrequent)) {
      result.push(i);
    }
  }
  return result;
}

function getStringFrequentArray(s: string): number[] {
  const frequentArray = new Array<number>(26).fill(0);
  for (const c of s) {
    frequentArray[charToNum(c)]!++;
  }
  return frequentArray;
}

function charToNum(c: string): number {
  return c.charCodeAt(0) - "a".charCodeAt(0);
}

function isEqualArray(a: number[], b: number[]): boolean {
  if (a.length !== b.length) {
    return false;
  }
  for (let i = 0; i < a.length; i++) {
    if (a[i] !== b[i]) {
      return false;
    }
  }
  return true;
}
