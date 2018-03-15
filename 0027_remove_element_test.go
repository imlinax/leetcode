package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func arrayNEqual(n int, a1, a2 []int) bool {
	if len(a1) < n || len(a2) < n {
		return false
	}
	for i := 0; i < n; i++ {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}

func Test_removeElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	expect := []int{2, 2}
	assert.Equal(t, len(expect), removeElement(nums, 3))
	assert.Equal(t, true, arrayNEqual(len(expect), expect, nums))
}
