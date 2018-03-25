/*
Given two binary strings, return their sum (also a binary string).

For example,
a = "11"
b = "1"
Return "100".
*/

package leetcode

import "fmt"

func addBinary(a string, b string) string {
	aIndex := len(a) - 1
	bIndex := len(b) - 1
	flag := 0
	result := ""
	for aIndex >= 0 && bIndex >= 0 {
		sum := int(a[aIndex]) - '0' + int(b[bIndex]) - '0' + flag
		if sum > 1 {
			flag = 1
			sum = sum - 2
		} else {
			flag = 0
		}
		result = fmt.Sprintf("%d", sum) + result
		aIndex--
		bIndex--
	}
	if aIndex < 0 {
		for bIndex >= 0 {
			sum := int(b[bIndex]) - int('0') + flag
			if sum > 1 {
				flag = 1
				sum = sum - 2
			} else {
				flag = 0
			}
			result = fmt.Sprintf("%d", sum) + result
			bIndex--
		}
	} else {
		for aIndex >= 0 {
			sum := int(a[aIndex]) - int('0') + flag
			if sum > 1 {
				flag = 1
				sum = sum - 2
			} else {
				flag = 0
			}
			result = fmt.Sprintf("%d", sum) + result
			aIndex--
		}
	}
	if flag == 1 {
		result = "1" + result
	}
	return result
}
