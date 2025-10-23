function closeStrings(word1: string, word2: string): boolean {
  const map1 = getCountMap(word1);
  const map2 = getCountMap(word2);
  return (
    [...map1.keys()].sort().toString() === [...map2.keys()].sort().toString() &&
    [...map1.values()].sort().toString() ===
      [...map2.values()].sort().toString()
  );
}

function getCountMap(word: string): Map<string, number> {
  const map = new Map<string, number>();
  for (const w of word) {
    if (map.has(w)) {
      map.set(w, map.get(w)! + 1);
    } else {
      map.set(w, 1);
    }
  }
  return map;
}
