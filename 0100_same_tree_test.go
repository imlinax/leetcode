package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isSameTree(t *testing.T) {
	assert.Equal(t, true, isSameTree(nil, nil))

	atree := &TreeNode{}
	btree := &TreeNode{}
	assert.Equal(t, true, isSameTree(atree, btree))

	atree = &TreeNode{}
	btree = nil
	assert.Equal(t, false, isSameTree(atree, btree))

	atree = &TreeNode{Val: 1}
	btree = &TreeNode{}
	assert.Equal(t, false, isSameTree(atree, btree))

	atree = &TreeNode{
		Val:  1,
		Left: &TreeNode{}}
	btree = &TreeNode{
		Val:  1,
		Left: &TreeNode{}}
	assert.Equal(t, true, isSameTree(atree, btree))
}
