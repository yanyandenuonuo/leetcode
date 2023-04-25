/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	swapIdx := 0
	for idx := 1; idx < len(nums); idx += 1 {
		if nums[swapIdx] == nums[idx] {
			continue
		}
		swapIdx += 1
		nums[swapIdx] = nums[idx]
	}
	return swapIdx + 1
}

// @lc code=end

