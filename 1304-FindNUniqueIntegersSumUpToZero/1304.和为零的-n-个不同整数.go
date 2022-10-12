/*
 * @lc app=leetcode.cn id=1304 lang=golang
 *
 * [1304] 和为零的 N 个不同整数
 */

// @lc code=start
func sumZero(n int) []int {
	res := make([]int, 0, n)
	for idx := 1; idx <= n/2; idx += 1 {
		res = append(res, idx)
		res = append(res, -idx)
	}
	if n&0x01 == 0x01 {
		res = append(res, 0)
	}
	return res
}

// @lc code=end

