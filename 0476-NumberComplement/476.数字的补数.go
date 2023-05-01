/*
 * @lc app=leetcode.cn id=476 lang=golang
 *
 * [476] 数字的补数
 */

// @lc code=start
func findComplement(num int) int {
	// solution1: 找出最高位的1，然后向低位进行逐位异或扫描操作
	// solution2: 找出最高位的1，然后构造1<<(idx+1)-1与num进行异或
	hightBit := 31
	for ; 1<<hightBit&num != 1<<hightBit; hightBit -= 1 {

	}
	return (1<<(hightBit+1) - 1) ^ num
}

// @lc code=end

