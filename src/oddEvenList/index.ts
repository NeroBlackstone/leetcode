class ListNode {
  val: number;
  next: ListNode | null;
  constructor(val?: number, next?: ListNode | null) {
    this.val = val === undefined ? 0 : val;
    this.next = next === undefined ? null : next;
  }
}
function oddEvenList(head: ListNode | null): ListNode | null {
  if (!head) {
    return null;
  }
  if (!head.next) {
    return head;
  }
  const evenNode = head.next;
  let oddNodeHead: ListNode | null = head;
  let evenNodeHead: ListNode | null = evenNode;
  while (oddNodeHead?.next && evenNodeHead?.next) {
    oddNodeHead!.next = evenNodeHead!.next;
    oddNodeHead = oddNodeHead!.next;
    evenNodeHead!.next = oddNodeHead?.next ?? null;
    evenNodeHead = evenNodeHead!.next;
  }
  oddNodeHead!.next = evenNode;
  return head;
}
