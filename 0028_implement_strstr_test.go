package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_strStr(t *testing.T) {
	haystack := "hello"
	needle = "ll"
	assert.Equal(t, 2, strStr(haystack, needle))

	haystack = "aaaaa"
	needle = "bba"
	assert.Equal(t, -1, strStr(haystack, needle))

}
