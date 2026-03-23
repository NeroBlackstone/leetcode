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

function isSymmetric(root: TreeNode | null): boolean {
  let result = true;
  if (!root) {
    return false;
  }
  retrieval(root.left, root.right);
  function retrieval(leftNode: TreeNode | null, rigthNode: TreeNode | null) {
    if ((!leftNode && rigthNode) || (!rigthNode && leftNode)) {
      result = false;
      return;
    } else if (leftNode && rigthNode) {
      if (leftNode.val !== rigthNode.val) {
        result = false;
        return;
      }
      retrieval(leftNode.left, rigthNode.right);
      retrieval(leftNode.right, rigthNode.left);
    } else {
      return;
    }
  }
  return result;
}
