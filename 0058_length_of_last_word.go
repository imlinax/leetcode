/*
Given a string s consists of upper/lower-case alphabets and empty space characters ' ', return the length of last word in the string.

If the last word does not exist, return 0.

Note: A word is defined as a character sequence consists of non-space characters only.

Example:

Input: "Hello World"
Output: 5
*/
package leetcode

func lengthOfLastWord(s string) int {
	space := rune(' ')
	count := 0
	flag := false
	for _, r := range s {
		if r == space {
			flag = true
		} else {
			if flag {
				count = 1
				flag = false
			} else {
				count++
			}
		}
	}
	return count
}
