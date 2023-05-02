/*
 * @lc app=leetcode.cn id=504 lang=golang
 *
 * [504] 七进制数
 */

// @lc code=start
func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	isNegative := false
	if num < 0 {
		num = -num
		isNegative = true
	}

	res := ""
	for ; num != 0; num /= 7 {
		res = string(rune('0'+num%7)) + res
	}

	if isNegative {
		res = "-" + res
	}
	return res
}

// @lc code=end

