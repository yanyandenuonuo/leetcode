/*
 * @lc app=leetcode.cn id=2438 lang=golang
 *
 * [2438] 二的幂数组中查询范围内的乘积
 */

// @lc code=start
func productQueries(n int, queries [][]int) []int {
	powers := make([]int, 0)
	for bitIdx := 0; n > 0; bitIdx += 1 {
		if n&0x01 == 0x01 {
			powers = append(powers, 0x01<<bitIdx)
		}
		n >>= 1
	}
	res := make([]int, 0)

	divisor := 1000000007
	for idx := 0; idx < len(queries); idx += 1 {
		powerInt := 1
		for idxStart := queries[idx][0]; idxStart <= queries[idx][1]; idxStart += 1 {
			powerInt *= powers[idxStart]
			powerInt %= divisor
		}
		res = append(res, powerInt)
	}

	return res
}

// @lc code=end

