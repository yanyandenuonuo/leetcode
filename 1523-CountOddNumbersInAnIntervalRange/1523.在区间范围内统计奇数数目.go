/*
 * @lc app=leetcode.cn id=1523 lang=golang
 *
 * [1523] 在区间范围内统计奇数数目
 */

// @lc code=start
func countOdds(low int, high int) int {
	// solution1: 枚举
	// solution2: 区间计算
	// return (high-low)/2 + (high-low)&0x01 + (low&0x01+high&0x01)/2

	// solution3: 前缀和：0-low有多少个奇数，0-high有多少个奇数
	return (high+1)/2 - (low)/2
}

// @lc code=end

