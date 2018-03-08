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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var dummyHead ListNode
	tail := &dummyHead

	pIndex1 := l1
	pIndex2 := l2
	flag := 0
	for pIndex1 != nil && pIndex2 != nil {
		sum := pIndex1.Val + pIndex2.Val + flag
		if sum >= 10 {
			flag = 1
			sum = sum % 10
		} else {
			flag = 0
		}
		newNode := &ListNode{
			Val:  sum,
			Next: nil}
		tail.Next = newNode
		tail = newNode
		pIndex1 = pIndex1.Next
		pIndex2 = pIndex2.Next
	}
	if pIndex1 == nil && pIndex2 == nil {
		if flag == 1 {
			tail.Next = &ListNode{
				Val:  1,
				Next: nil}
		}
	} else if pIndex1 == nil {
		for pIndex2 != nil {
			sum := pIndex2.Val + flag
			if sum >= 10 {
				flag = 1
				sum = sum % 10
			} else {
				flag = 0
			}

			tail.Next = &ListNode{
				Val:  sum,
				Next: nil}

			tail = tail.Next
			pIndex2 = pIndex2.Next
		}
		if flag == 1 {
			tail.Next = &ListNode{
				Val:  1,
				Next: nil}
		}
	} else {
		for pIndex1 != nil {
			sum := pIndex1.Val + flag
			if sum >= 10 {
				flag = 1
				sum = sum % 10
			} else {
				flag = 0
			}

			tail.Next = &ListNode{
				Val:  sum,
				Next: nil}

			tail = tail.Next
			pIndex1 = pIndex1.Next
		}
		if flag == 1 {
			tail.Next = &ListNode{
				Val:  1,
				Next: nil}
		}

	}

	return dummyHead.Next
}
