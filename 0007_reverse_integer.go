package leetcode

import (
	"math"
)

/*
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:

Input: 123
Output:  321
Example 2:

Input: -123
Output: -321
Example 3:

Input: 120
Output: 21
Note:
Assume we are dealing with an environment which could only hold integers within the 32-bit signed integer range. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
*/

func reverse(x int) int {
	var y, result int
	flag := 1
	if x < 0 {
		flag = -1
		x = int(math.Abs(float64(x)))
	}
	for x != 0 {
		y = x % 10
		result = result*10 + y
		x = x / 10
	}

	result = result * flag
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	return result
}
