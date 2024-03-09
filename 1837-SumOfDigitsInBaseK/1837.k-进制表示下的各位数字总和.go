/*
 * @lc app=leetcode.cn id=1837 lang=golang
 *
 * [1837] K 进制表示下的各位数字总和
 */

// @lc code=start
func sumBase(n int, k int) int {
	sum := 0
	for n > 0 {
		sum += n % k
		n /= k
	}

	return sum
}

// @lc code=end

