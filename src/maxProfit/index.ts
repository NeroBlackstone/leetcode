// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/
function maxProfit(prices: number[]): number {
  let minPrice = Number.MAX_VALUE;
  let profit = 0;
  for (const p of prices) {
    if (p < minPrice) {
      minPrice = p;
      // 注意这里不要让profit归零，即使不是从最低点买入的，也有利润
    } else {
      const possibleProfit = p - minPrice;
      if (possibleProfit > profit) {
        profit = possibleProfit;
      }
    }
  }
  return profit;
}
