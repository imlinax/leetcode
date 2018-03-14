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
func Test_removeDuplicates(t *testing.T) {
	input := []int{1, 1, 2}
	expect := []int{1, 2}
	assert.Equal(t, len(expect), removeDuplicates(input))
	assert.Equal(t, true, arrayNEqual(len(expect), input, expect))

	input = []int{1, 1, 1, 3, 4, 2, 2, 2, 2}
	expect = []int{1, 3, 4, 2}
	assert.Equal(t, len(expect), removeDuplicates(input))
	assert.Equal(t, true, arrayNEqual(len(expect), input, expect))
}
