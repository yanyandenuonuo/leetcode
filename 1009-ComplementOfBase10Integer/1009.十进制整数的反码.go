/*
 * @lc app=leetcode.cn id=1009 lang=golang
 *
 * [1009] 十进制整数的反码
 */

// @lc code=start
func bitwiseComplement(n int) int {
	// solution1: 找出最高位的1，然后向低位进行逐位异或扫描操作
	// solution2: 找出最高位的1，然后构造1<<(idx+1)-1与num进行异或
	if n == 0 {
		return 1
	}

	hightBit := 31
	for ; 1<<hightBit&n != 1<<hightBit; hightBit -= 1 {

	}
	return (1<<(hightBit+1) - 1) ^ n
}

// @lc code=end

