class ListNode {
  val: number;
  next: ListNode | null;
  constructor(val?: number, next?: ListNode | null) {
    this.val = val === undefined ? 0 : val;
    this.next = next === undefined ? null : next;
  }
}
// 输入：head = [1,2,3,4slow,3,2,1fast]
// null 4 3 2 1
// 输出：true
function isPalindrome(head: ListNode | null): boolean {
  if (!head || !head.next) {
    true;
  }
  let slow: ListNode | null | undefined = head;
  let fast: ListNode | null | undefined = head;
  // slow到中点了
  while (fast?.next) {
    slow = slow?.next;
    fast = fast.next.next;
  }
  let prevNode: ListNode | null | undefined = null;
  let currNode: ListNode | null | undefined = slow;
  while (currNode) {
    const temp = currNode.next;
    currNode.next = prevNode;
    prevNode = currNode;
    currNode = temp;
  }
  let left = head;
  while (prevNode) {
    if (prevNode.val !== left!.val) {
      return false;
    }
    prevNode = prevNode.next;
    left = left!.next;
  }
  return true;
}
