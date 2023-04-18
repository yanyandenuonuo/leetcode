/*
 * @lc app=leetcode.cn id=238 lang=golang
 *
 * [238] 除自身以外数组的乘积
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))
	answer[0] = 1

	// leftProduct
	for idx := 1; idx < len(nums); idx += 1 {
		answer[idx] = answer[idx-1] * nums[idx-1]
	}

	// rightProduct
	rightProduct := 1
	for idx := len(nums) - 1; idx >= 0; idx -= 1 {
		answer[idx] *= rightProduct
		rightProduct *= nums[idx]
	}

	return answer
}

// @lc code=end

