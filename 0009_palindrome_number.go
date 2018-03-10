/*
Determine whether an integer is a palindrome. Do this without extra space.
*/
package leetcode

func digitLength(x int) int {
	if x < 0 {
		x = x * -1
	}
	len := 1
	for x > 9 {
		x = x / 10
		len = len + 1
	}
	return len
}
func intPow(x, y int) int {
	result := 1
	for y > 0 {
		result = result * x
		y = y - 1
	}
	return result
}
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	l := digitLength(x)
	for i := 0; l-i-1 > i; i++ {
		if (x / intPow(10, l-i-1) % 10) != (x / intPow(10, i) % 10) {
			return false
		}

	}
	return true
}

func isPalindromeAnother(x int) bool {
	sum := 0
	origin := x
	if x < 0 {
		return false
	}
	for x > 0 {
		sum = sum*10 + x%10
		x = x / 10
	}
	return origin == sum
}
