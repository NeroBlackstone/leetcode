//https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
function maxProfit(prices: number[]): number {
  let profit = 0;
  for (let i = 0; i < prices.length - 1; i++) {
    const nowPrice = prices[i]!;
    const nextPrice = prices[i + 1]!;
    if (nowPrice < nextPrice) {
      profit += nextPrice - nowPrice;
    }
  }
  return profit;
}
