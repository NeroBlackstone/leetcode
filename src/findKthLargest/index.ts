import { MaxPriorityQueue } from "@datastructures-js/priority-queue";
console.log("leetcode");

// const pq = new MinPriorityQueue<number>();

// pq.enqueue(9);          // 插入
// pq.dequeue();           // 删除并返回最小元素
// pq.front();             // 查看最小元素
// pq.size();              // 大小
// pq.isEmpty();           // 是否为空

function findKthLargest(nums: number[], k: number): number {
  const pq = new MaxPriorityQueue<number>();
  for (const n of nums) {
    pq.enqueue(n);
  }
  let dequeueNum = null;
  for (let i = 0; i < k; i++) {
    dequeueNum = pq.dequeue();
  }
  return dequeueNum;
}
