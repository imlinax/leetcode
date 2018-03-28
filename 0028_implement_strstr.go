/*
Implement strStr().

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:

Input: haystack = "hello", needle = "ll"
Output: 2
Example 2:

Input: haystack = "aaaaa", needle = "bba"
Output: -1
*/
package leetcode

func strStr(haystack string, needle string) int {
	if len(needle) < 1 {
		return 0
	}
	for i := range haystack {
		flag := true
		if haystack[i] == needle[0] {
			for j := range needle {
				hIndex := i + j
				if hIndex > len(haystack)-1 {
					return -1
				}
				if needle[j] != haystack[hIndex] {
					flag = false
					break
				}
			}
			if flag {
				return i
			}
		}
	}
	return -1
}
