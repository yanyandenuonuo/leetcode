/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 */

// @lc code=start
func moveZeroes(nums []int)  {
    zeroIdx := -1
	for idx := 0; idx < len(nums); idx += 1 {
		if nums[idx] == 0 && zeroIdx == -1 {
			zeroIdx = idx
		} else if nums[idx] != 0 && zeroIdx >= 0 && zeroIdx != idx {
			nums[zeroIdx] = nums[idx]
			nums[idx] = 0
			if nums[zeroIdx+1] == 0 {
				zeroIdx += 1
			} else {
				zeroIdx = idx
			}
		}
	}
}
// @lc code=end

