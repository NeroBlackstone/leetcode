package climbStairs

import "math"

// https://leetcode.com/problems/climbing-stairs/
func climbStairs(n int) int {
	sqrt5 := math.Sqrt(5)
	pow := float64(n + 1)
	fibn := math.Pow((1+sqrt5)/2, pow) - math.Pow((1-sqrt5)/2, pow)
	return int(math.Round(fibn / sqrt5))
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Climbing Stairs.
//Memory Usage: 2 MB, less than 56.79% of Go online submissions for Climbing Stairs.
// Use Fibonacci Formula
