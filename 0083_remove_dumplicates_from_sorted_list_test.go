package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getListLen(head *ListNode) int {
	pIndex := head
	Length := 0
	for pIndex != nil {
		Length++
		pIndex = pIndex.Next
	}
	return Length
}

func Test_deleteDuplicates(t *testing.T) {
	list := &ListNode{}
	list = deleteDuplicates(list)
	assert.Equal(t, 1, getListLen(list))

	list = &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1}}
	list = deleteDuplicates(list)
	assert.Equal(t, 2, getListLen(list))

	list = &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 0}}
	list = deleteDuplicates(list)
	assert.Equal(t, 1, getListLen(list))

	list = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1}}}
	list = deleteDuplicates(list)
	assert.Equal(t, 1, getListLen(list))
}
