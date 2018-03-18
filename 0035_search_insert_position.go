/*
Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You may assume no duplicates in the array.

Example 1:

Input: [1,3,5,6], 5
Output: 2
Example 2:

Input: [1,3,5,6], 2
Output: 1
Example 3:

Input: [1,3,5,6], 7
Output: 4
Example 1:

Input: [1,3,5,6], 0
Output: 0
*/

package leetcode

func search(nums []int, left, right, target int) int {
	if left >= right {
		return left
	}
	mid := left + (right-left)/2

	if nums[mid] == target {
		return mid
	} else if nums[mid] > target {
		return search(nums, left, mid-1, target)
	}

	return search(nums, mid+1, right, target)

}
func searchInsert(nums []int, target int) int {
	index := search(nums, 0, len(nums)-1, target)
	if nums[index] < target {
		index = index + 1
	}

	return index
}
