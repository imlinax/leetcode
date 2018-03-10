package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_digitLength(t *testing.T) {
	assert.Equal(t, 1, digitLength(0))
	assert.Equal(t, 1, digitLength(1))
	assert.Equal(t, 2, digitLength(10))
	assert.Equal(t, 2, digitLength(55))
	assert.Equal(t, 3, digitLength(321))
	assert.Equal(t, 1, digitLength(-1))
	assert.Equal(t, 2, digitLength(-11))
}

func Test_isPalindrome(t *testing.T) {
	assert.Equal(t, true, isPalindrome(11))
	assert.Equal(t, true, isPalindrome(1))
	assert.Equal(t, false, isPalindrome(7788))
	assert.Equal(t, true, isPalindrome(787))
	assert.Equal(t, false, isPalindrome(21))
	assert.Equal(t, true, isPalindrome(78887))
}
