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

function longestZigZag(root: TreeNode | null): number {
  let max = 0;
  function dfs(node: TreeNode | null): [number, number] {
    if (!node) {
      return [-1, -1];
    }
    const [leftNodeLeftLength, leftNodeRightLength] = dfs(node.left);
    const [rightNodeLeftLength, rightNodeRightLength] = dfs(node.right);
    const leftLength = 1 + leftNodeRightLength;
    const rightLength = 1 + rightNodeLeftLength;
    max = Math.max(max, leftLength, rightLength);
    return [leftLength, rightLength];
  }
  dfs(root);
  return max;
}

// // 遍历树中所有节点,返回以node为根的子数里最大的zigZag值
// // 优化方向：知道了zigZag（node.left,"RIGHT"），就可以知道zigZag(node,"LEFT")+1
// function dfs(node: TreeNode | null): number {
//   if (node) {
//     // 左右方向的zigZag长度
//     const leftLength = zigZag(node, "LEFT");
//     const rightLength = zigZag(node, "RIGHT");
//     // 左右子节点的zigZag长度
//     const leftNodeLength = dfs(node.left);
//     const rightNodeLength = dfs(node.right);
//     return Math.max(leftLength, rightLength, leftNodeLength, rightNodeLength);
//   } else {
//     return 0;
//   }
// }

// type Direct = "LEFT" | "RIGHT";

// // 对某个节点执行zigZag,并且记录zigZag长度。
// function zigZag(node: TreeNode | null, direct: Direct): number {
//   if (!node) {
//     return 0;
//   }
//   let currentNode = node;
//   let length = 0;
//   // 循环的停止条件
//   while (currentNode) {
//     if (direct === "LEFT") {
//       currentNode = currentNode?.left;
//       if (currentNode) {
//         direct = "RIGHT";
//         length++;
//       }
//     } else {
//       currentNode = currentNode?.right;
//       if (currentNode) {
//         direct = "LEFT";
//         length++;
//       }
//     }
//   }
//   return length;
// }
