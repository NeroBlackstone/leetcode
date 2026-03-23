/**
 * Definition for singly-linked list.
 * class ListNode {
 *     val: number
 *     next: ListNode | null
 *     constructor(val?: number, next?: ListNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.next = (next===undefined ? null : next)
 *     }
 * }
 */

class ListNode {
  val: number;
  next: ListNode | null;
  constructor(val?: number, next?: ListNode | null) {
    this.val = val === undefined ? 0 : val;
    this.next = next === undefined ? null : next;
  }
}

function getIntersectionNode(
  headA: ListNode | null,
  headB: ListNode | null,
): ListNode | null {
  const set = new Set<ListNode>();
  let curNode = headA;
  while (curNode) {
    set.add(curNode);
    curNode = curNode.next;
  }
  curNode = headB;
  while (curNode) {
    if (set.has(curNode)) {
      return curNode;
    }
    curNode = curNode.next;
  }
  return null;
}
