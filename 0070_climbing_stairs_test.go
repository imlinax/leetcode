package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_climbStairs(t *testing.T) {
	assert.Equal(t, 2, climbStairs(2))
	assert.Equal(t, 3, climbStairs(3))
	assert.Equal(t, 89, climbStairs(10))
}
