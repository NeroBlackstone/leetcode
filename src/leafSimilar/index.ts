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

function leafSimilar(root1: TreeNode | null, root2: TreeNode | null): boolean {
  // 记录两个数的叶子结点
  const root1Leaf: number[] = [];
  const root2Leaf: number[] = [];
  // 查找第一颗树的叶子结点
  dfs(root1, root1Leaf);
  // 查找第二颗树的叶子结点
  dfs(root2, root2Leaf);
  return root1Leaf.toString() === root2Leaf.toString();
}

// 深度优先遍历
// 寻找树的所有叶子结点
function dfs(root: TreeNode | null, arr: number[]) {
  // 只有该节点的left和right都为空的时候
  if (!root!.left && !root!.right) {
    // 添加进数组
    return arr.push(root!.val);
  }
  // 判断左子树是否为空 传入左子树
  if (root!.left) dfs(root!.left, arr);
  // 判断右子树是否为空 传入右子树
  if (root!.right) dfs(root!.right, arr);
}
