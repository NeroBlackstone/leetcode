import { Stack } from "@datastructures-js/stack";

function decodeString(s: string): string {
  const stack = new Stack<string>();
  for (const c of s) {
    if (isDigit(c) || c === "[" || isLetter(c)) {
      stack.push(c);
    } else {
      const strArr = [];
      let popChar = stack.pop();
      while (popChar !== "[") {
        strArr.push(popChar);
        popChar = stack.pop();
      }
      const str = strArr.reverse().join("");
      popChar = stack.peek();
      let popDig = [];
      while (popChar && isDigit(popChar)) {
        popDig.push(stack.pop());
        popChar = stack.peek();
      }
      const digStr = popDig.reverse().join("");
      const times = digStr ? parseInt(digStr) : 1;
      for (const c of str.repeat(times)) {
        stack.push(c);
      }
    }
  }
  return stack.toArray().join("");
}

const isDigit = (ch: string) => /^[0-9]$/.test(ch);
const isLetter = (ch: string) => /^[A-Za-z]$/.test(ch);
