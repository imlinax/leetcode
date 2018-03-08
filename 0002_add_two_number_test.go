package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func listEqual(l1, l2 *ListNode) bool {
	p1 := l1
	p2 := l2
	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	if p1 != nil || p2 != nil {
		return false
	}
	return true
}
func listPrint(list *ListNode) {
	pIndex := list
	for pIndex != nil {
		fmt.Println(pIndex.Val)
		pIndex = pIndex.Next
	}
}

func arrayToList(array []int) *ListNode {
	var dummyNode ListNode
	tail := &dummyNode
	for _, val := range array {
		tail.Next = &ListNode{
			Val:  val,
			Next: nil}
		tail = tail.Next
	}
	return dummyNode.Next
}
func Test_addTwoNumbers(t *testing.T) {
	l1 := arrayToList([]int{2, 4, 3})
	l2 := arrayToList([]int{5, 6, 4})
	lExpect := arrayToList([]int{7, 0, 8})
	lResult := addTwoNumbers(l1, l2)

	assert.Equal(t, true, listEqual(lExpect, lResult))

	l1 = arrayToList([]int{1, 4, 3})
	l2 = arrayToList([]int{5, 1})
	lExpect = arrayToList([]int{6, 5, 3})
	lResult = addTwoNumbers(l1, l2)

	assert.Equal(t, true, listEqual(lExpect, lResult))

	l1 = arrayToList([]int{5})
	l2 = arrayToList([]int{5})
	lExpect = arrayToList([]int{0, 1})
	lResult = addTwoNumbers(l1, l2)

	assert.Equal(t, true, listEqual(lExpect, lResult))
}
