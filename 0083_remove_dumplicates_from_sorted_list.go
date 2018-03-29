/*
Given a sorted linked list, delete all duplicates such that each element appear only once.

For example,
Given 1->1->2, return 1->2.
Given 1->1->2->3->3, return 1->2->3.
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	pIndex := head
	for pIndex != nil {
		if pIndex.Next != nil {
			if pIndex.Next.Val == pIndex.Val {
				pIndex.Next = pIndex.Next.Next
			} else {
				pIndex = pIndex.Next
			}
		} else {
			return head
		}
	}
	return head
}
