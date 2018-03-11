/*
Write a function to find the longest common prefix string amongst an array of strings.
*/

package leetcode

import (
	"math"
)

func longestCommonPrefix(strs []string) string {
	var result []byte
	shortestLength := math.MaxInt64
	if len(strs) == 0 {
		return string(result)
	} else if len(strs) == 1 {
		return strs[0]
	}

	var shortestStr string
	var shortestIndex int
	for i, str := range strs {
		if len(str) < shortestLength {
			shortestLength = len(str)
			shortestStr = str
			shortestIndex = i
		}
	}
	for i, b := range []byte(shortestStr) {
		for si, str := range strs {
			if si == shortestIndex {
				continue
			}
			if str[i] != b {
				return string(result)
			}
		}
		result = append(result, b)
	}
	return string(result)
}
