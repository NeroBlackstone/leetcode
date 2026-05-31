function searchMatrix(matrix: number[][], target: number): boolean {
  for (const m of matrix) {
    if (m.includes(target)) {
      return true;
    }
  }
  return false;
}
