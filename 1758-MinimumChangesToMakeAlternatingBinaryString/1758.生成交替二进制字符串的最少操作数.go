/*
 * @lc app=leetcode.cn id=1758 lang=golang
 *
 * [1758] 生成交替二进制字符串的最少操作数
 */

// @lc code=start
func minOperations(s string) int {
	count0, count1 := 0, 0
	for idx, runeChar := range s {
		if idx&0x01 == int(runeChar-'0') {
			// 奇数位实际为1，但是期待基数为0
			count0 += 1
		} else {
			count1 += 1
		}
	}
	if count0 < count1 {
		return count0
	}
	return count1
}

// @lc code=end

