// 输入：s = "aab"
// 输出：[["a","a","b"],["aa","b"]]
function partition(s: string): string[][] {
  const result: string[][] = [];

  function backtrack(start: number, path: string[]) {
    if (start === s.length) {
      result.push([...path]);
      return;
    }
    for (let end = start + 1; end <= s.length; end++) {
      const subStr = s.slice(start, end);
      if (isPalindrome(subStr)) {
        path.push(subStr);
        backtrack(end, path);
        path.pop();
      }
    }
  }
  backtrack(0, []);
  return result;
}

function isPalindrome(str: string): boolean {
  let i = 0;
  while (i < str.length / 2) {
    if (str[i] !== str[str.length - 1 - i]) {
      return false;
    }
    i++;
  }
  return true;
}
