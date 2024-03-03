/*
 * @lc app=leetcode.cn id=2443 lang=golang
 *
 * [2443] 反转之后的数字和
 */

// @lc code=start
func sumOfNumberAndReverse(num int) bool {
	for idx := num / 2; idx <= num; idx += 1 {
		if idx+reverse(idx) == num {
			return true
		}
	}
	return false
}

func reverse(x int) int {
	res := 0
	for x != 0 {
		res = res*10 + x%10
		x /= 10
	}

	return res
}

// @lc code=end

