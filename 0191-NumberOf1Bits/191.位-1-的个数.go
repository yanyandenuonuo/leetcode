/*
 * @lc app=leetcode.cn id=191 lang=golang
 *
 * [191] 位1的个数
 */

// @lc code=start
func hammingWeight(num uint32) int {
	// solution1: 从低位到高位遍历
	/**
	bitCount := 0
	for num > 0 {
		bitCount += int(num & 0x01)
		num >>= 1
	}
	return bitCount
	*/

	// solution2: 每次变更最低的1
	bitCount := 0
	for num > 0 {
		num &= num - 1
		bitCount += 1

	}
	return bitCount

}

// @lc code=end

