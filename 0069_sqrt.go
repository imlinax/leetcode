/*
Implement int sqrt(int x).

Compute and return the square root of x.

x is guaranteed to be a non-negative integer.


Example 1:

Input: 4
Output: 2
Example 2:

Input: 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since we want to return an integer, the decimal part will be truncated.
*/
package leetcode

func mySqrt(x int) int {
	r := 0
	for {
		if r*r >= x {
			if r*r == x {
				return r
			} else {
				return r - 1
			}
		}
		r++
	}
	return r
}
