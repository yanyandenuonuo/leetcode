/*
 * @lc app=leetcode.cn id=405 lang=golang
 *
 * [405] 数字转换为十六进制数
 */

// @lc code=start
func toHex(num int) string {
	// solution: 每4bit转化为1位hex，从高位到低位依次转，忽略前导0
	if num == 0 {
		return "0"
	}

	res := make([]byte, 0, 8)
	for idx := 7; idx >= 0; idx -= 1 {
		hexNum := num >> (4 * idx) & 0x0f

		if hexNum == 0 && len(res) == 0 {
			continue
		}

		if hexNum < 10 {
			res = append(res, byte('0'+hexNum))
		} else {
			res = append(res, byte('a'+hexNum-10))
		}
	}
	return string(res)
}

// @lc code=end

