/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 */

// @lc code=start
func firstMissingPositive(nums []int) int {
	// solution1: 排除负值，再用负值占位置
	/**
	for idx, val := range nums {
		// 排除负值
		if val <= 0 {
			nums[idx] = len(nums) + 1
		}
	}
	// 用负值占位置，将nums[idx]-1做为索引，将对应位置的数字设置为负数
	for _, val := range nums {
		if val > len(nums) || -val > len(nums) {
			// not in [1:n]
			continue
		}
		if val < 0 {
			// idx+1所对应的数字已经被处理过
			val = -val
		}
		if nums[val-1] < 0 {
			continue
		}
		nums[val-1] = -nums[val-1]
	}
	for idx := 0; idx < len(nums); idx += 1 {
		if nums[idx] > 0 {
			return idx + 1
		}
	}
	return len(nums) + 1
	*/

	// solution2: 将[1:n]的数字放到符合条件的位置上
	for idx := 0; idx < len(nums); idx += 1 {
		for val := nums[idx]; val > 0 && val <= len(nums) && nums[val-1] != val; val = nums[idx] {
			nums[val-1], nums[idx] = val, nums[val-1]
		}
	}
	for idx, val := range nums {
		if val != idx+1 {
			return idx + 1
		}
	}
	return len(nums) + 1
}

// @lc code=end

