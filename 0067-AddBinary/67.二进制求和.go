/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 */

// @lc code=start
func addBinary(a string, b string) string {
	if len(a) > len(b) {
		a, b = b, a
	}

	carry := uint8(0)
	res := ""

	for offset := 0; offset < len(b); offset += 1 {
		aBit := uint8(0)
		if offset < len(a) {
			aBit = a[len(a)-1-offset] - '0'
		}

		bBit := b[len(b)-1-offset] - '0'

		bitSum := aBit + bBit + carry

		res = string('0'+bitSum&0x01) + res
		if bitSum&0x02 == 0x02 {
			carry = 1
		} else {
			carry = 0
		}
	}

	// 补进位
	if carry == 1 {
		res = string('0'+carry) + res
	}
	return res
}

// @lc code=end

