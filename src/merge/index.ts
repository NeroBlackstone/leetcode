// 示例 1：

// 输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
// 输出：[[1,6],[8,10],[15,18]]
// 解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
function merge(intervals: number[][]): number[][] {
  if (intervals.length <= 1) {
    return intervals;
  }
  const result: number[][] = [];
  intervals.sort((a, b) => a[0]! - b[0]!);
  let [start, end] = intervals[0]!;
  for (let i = 1; i < intervals.length; i++) {
    const [curStart, curEnd] = intervals[i]!;
    if (curStart! <= end!) {
      end = Math.max(end!, curEnd!);
    } else {
      result.push([start!, end!]);
      start = curStart;
      end = curEnd;
    }
  }
  result.push([start!, end!]);
  return result;
}
