/*
 * @lc app=leetcode.cn id=1486 lang=golang
 *
 * [1486] 数组异或操作
 */

// @lc code=start
func xorOperation(n int, start int) int {
	// solution1: 逐步操作
	res := 0
	for idx := 0; idx < n; idx += 1 {
		res ^= (start + 2*idx)
	}
	return res

	// solution2: 数序推演，以下仅做参考。
	/**
	sumXor := func(x int) int {
		switch x % 4 {
		case 0:
			return x
		case 1:
			return 1
		case 2:
			return x + 1
		default:
			return 0
		}
	}

	s, e := start>>1, n&start&1
	res := sumXor(s-1) ^ sumXor(s+n-1)
	return res<<1 | e
	*/
}

// @lc code=end

