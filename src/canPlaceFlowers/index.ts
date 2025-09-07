function canPlaceFlowers(flowerbed: number[], n: number): boolean {
  if (n === 0) {
    return true;
  }
  let lastPlaced = -2;
  for (let i = 0; i < flowerbed.length; i++) {
    if (flowerbed[i] === 1) {
      const canPlaceCount = i - lastPlaced - 1;
      if (canPlaceCount > 1) {
        n = n - (Math.ceil(canPlaceCount / 2) - 1);
        if (n <= 0) {
          return true;
        }
      }
      lastPlaced = i;
    }
  }
  const lastCount = flowerbed.length - 1 - lastPlaced;
  n = n - Math.floor(lastCount / 2);
  if (n <= 0) {
    return true;
  }
  return false;
}
// https://leetcode.com/problems/can-place-flowers/description/
