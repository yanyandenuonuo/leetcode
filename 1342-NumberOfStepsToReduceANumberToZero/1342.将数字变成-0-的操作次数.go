/*
 * @lc app=leetcode.cn id=1342 lang=golang
 *
 * [1342] 将数字变成 0 的操作次数
 */

// @lc code=start
func numberOfSteps(num int) int {
	res := 0
	for num > 0 {
		if num%2 == 0 {
			num = num >> 1
		} else {
			num -= 1
		}
		res += 1
	}
	return res
}

// @lc code=end

