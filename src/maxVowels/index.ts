function maxVowels(s: string, k: number): number {
  let current = [...s.slice(0, k)].reduce((acc, b) => acc + isVolwel(b), 0);
  let max = current;
  for (let i = 1; i < s.length - k + 1; i++) {
    current = current - isVolwel(s[i - 1]!) + isVolwel(s[i + k - 1]!);
    max = Math.max(current, max);
  }
  return max;
}

function isVolwel(s: string): number {
  return s === "a" || s === "e" || s === "i" || s === "o" || s === "u" ? 1 : 0;
}
