package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lengthOfLastWord(t *testing.T) {
	assert.Equal(t, 5, lengthOfLastWord("Hello World"))
	assert.Equal(t, 5, lengthOfLastWord("Hello"))
	assert.Equal(t, 5, lengthOfLastWord(" Hello"))
	assert.Equal(t, 5, lengthOfLastWord("Hello "))
}
