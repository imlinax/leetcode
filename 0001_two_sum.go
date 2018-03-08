package leetcode

/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int, len(nums))
	for i, v := range nums {
		numMap[v] = i
	}
	for index1, v := range nums {
		res := target - v
		index2, ok := numMap[res]
		if ok && index1 != index2 {
			return []int{index1, index2}
		}
	}
	return nil
}
