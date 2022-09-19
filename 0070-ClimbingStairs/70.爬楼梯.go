/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	sumArray := make([]int, n+1)
	sumArray[0] = 1
	sumArray[1] = 1
	for idx := 2; idx <= n; idx += 1 {
		sumArray[idx] = sumArray[idx-1] + sumArray[idx-2]
	}
	return sumArray[n]
}

// @lc code=end

