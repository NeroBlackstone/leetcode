function generateParenthesis(n: number): string[] {
  const result: string[] = [];

  function backtrack(path: string, leftUse: number, rightUse: number) {
    if (path.length === n * 2) {
      result.push(path);
      return;
    }

    if (leftUse < n) {
      backtrack(`${path}(`, leftUse + 1, rightUse);
    }
    if (rightUse < leftUse) {
      backtrack(`${path})`, leftUse, rightUse + 1);
    }
  }
  backtrack("", 0, 0);
  return result;
}
