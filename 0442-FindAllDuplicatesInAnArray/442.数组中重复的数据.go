/*
 * @lc app=leetcode.cn id=442 lang=golang
 *
 * [442] 数组中重复的数据
 */

// @lc code=start
func findDuplicates(nums []int) []int {
	for _, val := range nums {
		if val < 0 {
			val = -val
		}

		idx := (val - 1) % len(nums)
		if nums[idx] > 0 {
			nums[idx] += len(nums)
		}
		nums[idx] = -nums[idx]
	}
	res := make([]int, 0, len(nums)/2)
	for idx, val := range nums {
		if val <= len(nums) {
			continue
		}

		res = append(res, idx+1)
	}
	return res
}

// @lc code=end

