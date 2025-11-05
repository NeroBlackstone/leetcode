import { Stack } from "@datastructures-js/stack";

function asteroidCollision(asteroids: number[]): number[] {
  const stack = new Stack<number>();
  for (const a of asteroids) {
    while (true) {
      if (stack.size() === 0 || !(stack.peek()! > 0 && a < 0)) {
        stack.push(a);
        break;
      } else if (stack.peek()! + a === 0) {
        stack.pop();
        break;
      } else if (stack.peek()! + a > 0) {
        break;
      } else {
        stack.pop();
      }
    }
  }
  return stack.toArray();
}
