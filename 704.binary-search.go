/*
 * @lc app=leetcode id=704 lang=golang
 *
 * [704] Binary Search
 */

// @lc code=start
func search(nums []int, target int) int {
	rightIdx := len(nums) - 1

	for leftIdx := 0; leftIdx <= rightIdx; {
		midIdx := leftIdx + (rightIdx - leftIdx) / 2
		if nums[midIdx] == target {
			return midIdx
		} else if nums[midIdx] > target {
			rightIdx = midIdx - 1
		} else if nums[midIdx] < target {
			leftIdx = midIdx + 1
		}
	}
	return -1
}
// @lc code=end

