/*
 * @lc app=leetcode.cn id=762 lang=golang
 *
 * [762] 二进制表示中质数个计算置位
 */

// @lc code=start
func countPrimeSetBits(left int, right int) int {
	res := 0
	for idx := left; idx <= right; idx += 1 {
		// 1 <= left <= right <= 10^6 < 2^20
		// 所以区间内的数字最多为19位，1最多出现19次，而20以内的质数只有2,3,5,7,11,13,17,19
		// 665772 = 1<<2 | 1<<3 | 1<<5 | 1<<7 | 1<<11 | 1<<13 | 1<<17 | 1<<19
		if 1<<bits.OnesCount(uint(idx))&665772 != 0 {
			res += 1
		}
	}
	return res
}

// @lc code=end

