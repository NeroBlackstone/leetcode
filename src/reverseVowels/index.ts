function reverseVowels(s: string): string {
  const arr = s.split("");
  const vowels = ["a", "e", "i", "o", "u"];
  let start = 0;
  let end = s.length - 1;
  while (start < end) {
    while (!vowels.includes(arr[start]!.toLowerCase()) && start < end) {
      start++;
    }
    while (!vowels.includes(arr[end]!.toLowerCase()) && start < end) {
      end--;
    }
    if (start < end) {
      const temp = s[end]!;
      arr[end] = s[start]!;
      arr[start] = temp;
    } else {
      break;
    }
    start++;
    end--;
  }
  return arr.join("");
}
