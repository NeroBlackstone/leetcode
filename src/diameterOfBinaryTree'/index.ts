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

function diameterOfBinaryTree(root: TreeNode | null): number {
  let diameter = 0;
  depth(root);
  function depth(root: TreeNode | null): number {
    if (!root) {
      return 0;
    }
    const leftDepth = depth(root.left);
    const rightDepth = depth(root.right);
    diameter = Math.max(diameter, leftDepth + rightDepth);
    return Math.max(leftDepth, rightDepth) + 1;
  }
  return diameter;
}
