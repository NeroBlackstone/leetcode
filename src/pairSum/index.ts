class ListNode {
  val: number;
  next: ListNode | null;
  constructor(val?: number, next?: ListNode | null) {
    this.val = val === undefined ? 0 : val;
    this.next = next === undefined ? null : next;
  }
}

function pairSum(head: ListNode | null): number {
  let slow = head;
  let fast = head!.next;
  while (fast?.next) {
    slow = slow!.next;
    fast = fast.next.next;
  }
  let tailNode = reverseList(slow!.next);
  let max = -Infinity;
  let headNode = head;
  while (tailNode) {
    max = Math.max(tailNode.val + headNode!.val, max);
    tailNode = tailNode.next;
    headNode = headNode!.next;
  }
  return max;
}
function reverseList(head: ListNode | null): ListNode | null {
  if (!head || !head.next) {
    return head;
  }
  let currNode = head;
  let nextNode: ListNode | null = head.next;
  head.next = null;
  while (nextNode) {
    const nextNextNode: ListNode | null = nextNode.next;
    nextNode.next = currNode;
    currNode = nextNode;
    nextNode = nextNextNode;
  }
  return currNode;
}
