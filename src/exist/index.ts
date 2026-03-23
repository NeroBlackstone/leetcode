function exist(board: string[][], word: string): boolean {
  const rows = board.length;
  const cols = board[0]!.length;
  let found = false;
  for (let x = 0; x < rows; x++) {
    for (let y = 0; y < cols; y++) {
      if (board[x]![y] === word[0]) {
        backtrack(x, y, 0);
      }
    }
  }

  function backtrack(x: number, y: number, index: number) {
    if (found) return;
    // 出格子
    if (x < 0 || x >= rows || y < 0 || y >= cols) {
      return;
    }
    // 已访问
    if (board[x]![y] === "#") {
      return;
    }
    let temp = "";
    // 字不对
    if (board[x]![y] !== word[index]) {
      return;
    } else {
      temp = board[x]![y]!;
      board[x]![y] = "#";
    }

    if (index === word.length - 1) {
      found = true;
      return;
    }

    backtrack(x + 1, y, index + 1);
    backtrack(x - 1, y, index + 1);
    backtrack(x, y + 1, index + 1);
    backtrack(x, y - 1, index + 1);

    board[x]![y] = temp;
  }
  return found;
}
