//https://leetcode.com/problems/rotate-image/
function rotate(matrix: number[][]): void {
  const len = matrix.length;
  //转置
  for (let i = 0; i < len; i++) {
    for (let j = i; j < len; j++) {
      [matrix[i]![j], matrix[j]![i]] = [matrix[j]![i]!, matrix[i]![j]!];
    }
  }
  for (const row of matrix) {
    row.reverse();
  }
}
