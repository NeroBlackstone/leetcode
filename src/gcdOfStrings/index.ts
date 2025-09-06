function gcdOfStrings(str1: string, str2: string): string {
  if (str1 + str2 !== str2 + str1) {
    return "";
  }
  return str1.slice(0, gcd(str1.length, str2.length));
}

function gcd(a: number, b: number): number {
  if (a < b) {
    [a, b] = [b, a];
  }
  let reminder = a % b;
  while (reminder !== 0) {
    a = b;
    b = reminder;
    reminder = a % b;
  }
  return b;
}
// https://leetcode.com/problems/greatest-common-divisor-of-strings/description/
