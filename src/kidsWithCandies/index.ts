function kidsWithCandies(candies: number[], extraCandies: number): boolean[] {
  const max = Math.max(...candies);
  return candies.map((e) => e + extraCandies >= max);
}
//https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/description/
