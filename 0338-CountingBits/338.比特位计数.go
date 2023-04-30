/*
 * @lc app=leetcode.cn id=338 lang=golang
 *
 * [338] 比特位计数
 */

// @lc code=start
func countBits(n int) []int {
	// solution1: 逐个统计1的位数
	/**
	res := make([]int, n+1)
	for idx := 1; idx <= n; idx += 1 {
		res[idx] = getNumBit(idx)
	}
	return res
	*/

	// solution2: 动态规划，最低有效位
	//			  bit[x] = bix[x/2] + x & 0x01
	/**
	res := make([]int, n+1)
	for idx := 1; idx <= n; idx += 1 {
		res[idx] = res[idx>>1] + idx&0x01
	}
	return res
	*/

	// solution3: 动态规划，最高有效位
	//			  bit[x] = bix[x-hightBit] + 1
	/**
	res := make([]int, n+1)
	hightBit := 0
	for idx := 1; idx <= n; idx += 1 {
		if idx&(idx-1) == 0 {
			hightBit = idx
		}

		res[idx] = res[idx-hightBit] + 1
	}
	return res
	*/

	// solution4: 动态规划，最低设置位
	//			  bit[x] = bix[x&(x-1)] + 1
	res := make([]int, n+1)
	for idx := 1; idx <= n; idx += 1 {
		res[idx] = res[idx&(idx-1)] + 1
	}
	return res
}

func getNumBit(num int) int {
	counter := 0
	for num > 0 {
		num &= num - 1
		counter += 1
	}
	return counter
}

// @lc code=end

