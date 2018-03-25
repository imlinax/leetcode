package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_plusOne(t *testing.T) {
	assert.Equal(t, []int{1}, plusOne([]int{0}))
	assert.Equal(t, []int{1, 0}, plusOne([]int{9}))
	assert.Equal(t, []int{1, 0, 0}, plusOne([]int{9, 9}))
	assert.Equal(t, []int{3}, plusOne([]int{2}))
	assert.Equal(t, []int{3, 2}, plusOne([]int{3, 1}))
}
