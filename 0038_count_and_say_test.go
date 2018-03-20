package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_countAndSay(t *testing.T) {
	assert.Equal(t, "1", countAndSay(1))
	assert.Equal(t, "1211", countAndSay(4))
	assert.Equal(t, "312211", countAndSay(6))
}
