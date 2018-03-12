package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var newHead ListNode
	newHead.Next = nil
	p1 := l1
	p2 := l2
	pNew := &newHead
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			next := p1.Next
			pNew.Next = p1
			p1 = next
		} else {
			next := p2.Next
			pNew.Next = p2
			p2 = next
		}
		pNew = pNew.Next
	}
	if p1 != nil {
		pNew = p1
	} else {
		pNew = p2
	}
	return newHead.Next
}
