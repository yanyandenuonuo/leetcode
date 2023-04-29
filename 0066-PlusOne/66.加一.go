/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	carry := 1
	for idx := len(digits) - 1; idx >= 0 && carry > 0; idx -= 1 {
		bitSum := digits[idx] + carry
		digits[idx] = bitSum % 10
		carry = bitSum / 10
	}

	if carry > 0 {
		return append([]int{carry}, digits...)
	}

	return digits
}

// @lc code=end

