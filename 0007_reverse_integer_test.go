package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverse(t *testing.T) {
	assert.Equal(t, 1, reverse(100))
	assert.Equal(t, -1, reverse(-100))
	assert.Equal(t, -23, reverse(-32))
	assert.Equal(t, 21, reverse(120))
	assert.Equal(t, 0, reverse(1534236469))
}
