/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	ans := float64(1)
	reverse := false
	if n < 0 {
		reverse = true
		n = -n
	}
	for n > 0 {
		if n&0x01 == 0x01 {
			ans *= x
		}
		n >>= 1
		x *= x
	}
	if reverse {
		return float64(1) / ans
	}
	return ans
}

// @lc code=end

