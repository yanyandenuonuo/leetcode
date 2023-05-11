/*
 * @lc app=leetcode.cn id=1342 lang=golang
 *
 * [1342] 将数字变成 0 的操作次数
 */

// @lc code=start
func numberOfSteps(num int) int {
	// solution: 模拟
	counter := 0
	for num > 0 {
		if num%2 == 0 {
			num = num >> 1
		} else {
			num -= 1
		}
		counter += 1
	}

	return counter
}

// @lc code=end

