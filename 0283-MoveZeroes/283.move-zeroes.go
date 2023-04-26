/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 */

// @lc code=start
func moveZeroes(nums []int) {
	// solution1: 直接移动非0到最前方，末尾用0覆盖
	// solution2: 双指针移动，逐个替换0的位置
	for zeroIdx, idx := 0, 0; zeroIdx <= idx && zeroIdx < len(nums) && idx < len(nums); {
		if nums[zeroIdx] != 0 {
			zeroIdx += 1
			if zeroIdx > idx {
				idx = zeroIdx
			}
			continue
		}
		if nums[idx] == 0 {
			idx += 1
			continue
		}
		nums[zeroIdx], nums[idx] = nums[idx], nums[zeroIdx]
		zeroIdx += 1
		idx += 1
	}
}

// @lc code=end

