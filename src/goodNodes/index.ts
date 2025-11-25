class TreeNode {
  val: number;
  left: TreeNode | null;
  right: TreeNode | null;
  constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
    this.val = val === undefined ? 0 : val;
    this.left = left === undefined ? null : left;
    this.right = right === undefined ? null : right;
  }
}

function goodNodes(root: TreeNode | null): number {
  return dfs(root, -Infinity);
}
function dfs(root: TreeNode | null, pathMax: number): number {
  if (!root) {
    return 0;
  }
  let res = 0;
  if (root.val >= pathMax) {
    res++;
    pathMax = root.val;
  }
  return res + dfs(root.left, pathMax) + dfs(root.right, pathMax);
}
