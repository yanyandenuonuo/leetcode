/*
 * @lc app=leetcode.cn id=2460 lang=golang
 *
 * [2460] 对数组执行操作
 */

// @lc code=start
func applyOperations(nums []int) []int {
	for idx := 0; idx < len(nums)-1; idx += 1 {
		if nums[idx] == 0 {
			continue
		}

		if nums[idx] == nums[idx+1] {
			nums[idx] = 2 * nums[idx+1]
			nums[idx+1] = 0
		}
	}

	for zeroIdx, idx := -1, 0; idx < len(nums); idx += 1 {
		if nums[idx] == 0 {
			if zeroIdx < 0 {
				zeroIdx = idx
			}

			continue
		}

		if zeroIdx >= 0 && zeroIdx < idx {
			nums[zeroIdx], nums[idx] = nums[idx], nums[zeroIdx]

			if nums[zeroIdx+1] == 0 {
				zeroIdx += 1
			}
		}
	}

	return nums
}

// @lc code=end

