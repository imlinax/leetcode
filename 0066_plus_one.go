/*
Given a non-negative integer represented as a non-empty array of digits, plus one to the integer.

You may assume the integer do not contain any leading zero, except the number 0 itself.

The digits are stored such that the most significant digit is at the head of the list.
*/

package leetcode

func plusOne(digits []int) []int {
	length := len(digits)
	result := make([]int, 0)
	flag := 0
	for i := length - 1; i >= 0; i-- {
		val := digits[i]
		sum := val + flag
		if i == length-1 {
			sum++
		}
		if sum > 9 {
			flag = 1
			result = append([]int{0}, result...)
		} else {
			flag = 0
			result = append([]int{sum}, result...)
		}
	}
	if flag == 1 {
		result = append([]int{1}, result...)
	}
	return result
}
