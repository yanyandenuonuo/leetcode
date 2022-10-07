/*
 * @lc app=leetcode.cn id=2427 lang=golang
 *
 * [2427] 公因子的数目
 */

// @lc code=start
func commonFactors(a int, b int) int {
	if a > b {
		a, b = b, a
	}
	total := 0
	for idx := 1; idx <= a; idx += 1 {
		if a%idx == 0 && b%idx == 0 {
			total += 1
		}
	}

	return total
}

// @lc code=end

