function mergeAlternately(word1: string, word2: string): string {
  const answer = [];
  const word1Len = word1.length;
  const word2Len = word2.length;
  let i = 0;
  while (i < Math.min(word1Len, word2Len)) {
    const nextChar1 = word1[i];
    const nextChar2 = word2[i];
    if (nextChar1) {
      answer.push(nextChar1);
    }
    if (nextChar2) {
      answer.push(nextChar2);
    }
    i++;
  }
  if (i === word1Len) {
    answer.push(word2.slice(i, word2Len));
  } else {
    answer.push(word1.slice(i, word1Len));
  }
  // 记住，join永远比直接+拼字符串快
  return answer.join("");
}
