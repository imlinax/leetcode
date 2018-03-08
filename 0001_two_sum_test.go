package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_twoSum(t *testing.T) {
	res := twoSum([]int{2, 7, 11, 15}, 9)
	assert.Equal(t, []int{0, 1}, res)

}
