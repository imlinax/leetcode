package leetcode

func removeElement(nums []int, val int) int {
	index := 0
	for i, v := range nums {
		if v == val {
			continue
		}
		nums[index] = nums[i]
		index++
	}
	return index
}
