/*
 * @lc app=leetcode.cn id=896 lang=golang
 *
 * [896] 单调数列
 */

// @lc code=start
func isMonotonic(nums []int) bool {
	// solution1: 分别遍历单调递增和单调递减
	// solution2: 一次遍历，使用flag标识
	inc, dec := true, true
	for idx := 0; idx < len(nums)-1; idx += 1 {
		if nums[idx] > nums[idx+1] {
			inc = false
		}

		if nums[idx] < nums[idx+1] {
			dec = false
		}
	}

	return inc || dec
}

// @lc code=end

