package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestCommonPrefix(t *testing.T) {
	assert.Equal(t, "abc", longestCommonPrefix([]string{"abc", "abc123"}))
	assert.Equal(t, "", longestCommonPrefix([]string{"abc", "123"}))
	assert.Equal(t, "abc", longestCommonPrefix([]string{"abc"}))
	assert.Equal(t, "", longestCommonPrefix([]string{"ca", "a"}))
}
