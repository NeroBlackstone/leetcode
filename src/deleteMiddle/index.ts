class ListNode {
  val: number;
  next: ListNode | null;
  constructor(val?: number, next?: ListNode | null) {
    this.val = val === undefined ? 0 : val;
    this.next = next === undefined ? null : next;
  }
}
function deleteMiddle(head: ListNode | null): ListNode | null {
  if (!head?.next) {
    return null;
  }
  let slow: ListNode | null = head;
  let fast: ListNode | null = head;
  let slowPre = null;
  while (fast?.next) {
    slowPre = slow;
    fast = fast.next.next;
    slow = slow!.next;
  }
  slowPre!.next = slow!.next;
  return head;
}
