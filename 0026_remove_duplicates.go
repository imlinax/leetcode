/*
Given a sorted array, remove the duplicates in-place such that each element appear only once and return the new length.

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

Example:

Given nums = [1,1,2],

Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively.
It doesn't matter what you leave beyond the new length.
*/
package leetcode

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length < 2 {
		return length
	}
	index := 0
	cvalue := nums[index]
	for i := 1; i < length; i++ {
		if nums[i] != cvalue {
			index++
			nums[index] = nums[i]
			cvalue = nums[i]
		}
	}
	return index + 1
}
