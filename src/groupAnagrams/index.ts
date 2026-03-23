function groupAnagrams(strs: string[]): string[][] {
  const map = new Map<string, string[]>();
  for (const s of strs) {
    const sortedString = s.split("").sort().join("");
    map.has(sortedString)
      ? map.get(sortedString)?.push(s)
      : map.set(sortedString, [s]);
  }
  return Array.from(map.values());
}
