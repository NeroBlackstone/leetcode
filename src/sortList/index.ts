class ListNode {
  val: number;
  next: ListNode | null;
  constructor(val?: number, next?: ListNode | null) {
    this.val = val === undefined ? 0 : val;
    this.next = next === undefined ? null : next;
  }
}

function sortList(head: ListNode | null): ListNode | null {
  if (!head || !head.next) {
    return head;
  }
  //快慢指针断链
  let slow: ListNode | null = head;
  let fast: ListNode | null = head;
  let preSlow: ListNode | null = head;
  while (fast?.next) {
    preSlow = slow;
    slow = slow!.next;
    fast = fast.next.next;
  }
  preSlow!.next = null;
  const left = sortList(head);
  const right = sortList(slow);
  return mergeList(left, right);
}

// 合并有序链表
function mergeList(
  left: ListNode | null,
  right: ListNode | null,
): ListNode | null {
  const head: ListNode = new ListNode(-Infinity, null);
  let cuur: ListNode = head;
  while (left && right) {
    if (left.val <= right.val) {
      cuur.next = left;
      left = left.next;
    } else {
      cuur.next = right;
      right = right.next;
    }
    cuur = cuur.next;
  }

  cuur.next = left || right;
  return head.next;
}
