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
	minLength := len(a)

	carry := uint8(0)
	res := ""
	offset := 0
	for ; offset < minLength; offset += 1 {
		aBit := a[len(a)-1-offset] - '0'
		bBit := b[len(b)-1-offset] - '0'
		bitNum := aBit + bBit + carry
		res = string('0'+bitNum&0x01) + res
		if bitNum&0x02 == 0x02 {
			carry = uint8(1)
		} else {
			carry = uint8(0)
		}
	}

	for ; carry == 1 && offset < len(b); offset += 1 {
		bBit := b[len(b)-1-offset] - '0'
		bitNum := bBit + carry
		res = string('0'+bitNum&0x01) + res
		if bitNum&0x02 == 0x02 {
			carry = uint8(1)
		} else {
			carry = uint8(0)
		}
	}

	// 补齐b字符高位
	if offset != len(b) {
		res = b[0:len(b)-offset] + res
	}

	// 补进位
	if carry == 1 {
		res = string('0'+carry) + res
	}
	return res
}

// @lc code=end

