/*
 * @lc app=leetcode.cn id=2210 lang=golang
 *
 * [2210] 统计数组中峰和谷的数量
 */

// @lc code=start
func countHillValley(nums []int) int {
	// solution1: 压缩数组后计数
	// solution2: 遍历
	counter := 0
	for preVal, idx := nums[0], 1; idx < len(nums)-1; {
		if preVal == nums[idx] {
			idx += 1
			continue
		}

		nextIdx := idx + 1
		for nums[nextIdx] == nums[idx] {
			nextIdx += 1

			if nextIdx == len(nums) {
				return counter
			}
		}

		if (preVal > nums[idx] && nums[nextIdx] > nums[idx]) ||
			(preVal < nums[idx] && nums[nextIdx] < nums[idx]) {
			counter += 1
		}

		preVal = nums[idx]
		idx = nextIdx
	}

	return counter
}

// @lc code=end

