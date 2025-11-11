class ListNode {
  val: number;
  next: ListNode | null;
  constructor(val?: number, next?: ListNode | null) {
    this.val = val === undefined ? 0 : val;
    this.next = next === undefined ? null : next;
  }
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
