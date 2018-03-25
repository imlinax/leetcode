package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_addBinary(t *testing.T) {
	assert.Equal(t, "100", addBinary("11", "1"))
	assert.Equal(t, "100", addBinary("1", "11"))
	assert.Equal(t, "10", addBinary("1", "1"))
	assert.Equal(t, "11", addBinary("10", "1"))
}
