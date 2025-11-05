import { Stack } from "@datastructures-js/stack";

function removeStars(s: string): string {
  const stack = new Stack<string>();
  for (const c of s) {
    if (c === "*") {
      stack.pop();
    } else {
      stack.push(c);
    }
  }
  return stack.toArray().join("");
}
