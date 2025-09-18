// https://leetcode.com/problems/string-compression/description/
function compress(chars: string[]): number {
  if (chars.length === 0) {
    return 0;
  }
  let currChar = chars[0]!;
  let charLength = 1;
  let changeIndex = 0;
  for (let i = 1; i < chars.length; i++) {
    const c = chars[i]!;
    if (c === currChar) {
      charLength++;
    } else {
      const subString =
        charLength === 1 ? currChar : currChar + charLength.toString();
      for (const s of subString) {
        chars[changeIndex] = s;
        changeIndex++;
      }
      charLength = 1;
      currChar = c;
    }
  }
  const subString =
    charLength === 1 ? currChar : currChar + charLength.toString();
  for (const s of subString) {
    chars[changeIndex] = s;
    changeIndex++;
  }
  chars.length = changeIndex;
  return chars.length;
}
