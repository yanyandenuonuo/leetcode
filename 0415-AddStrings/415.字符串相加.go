/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 */

// @lc code=start
func addStrings(num1 string, num2 string) string {
	if len(num1) > len(num2) {
		num1, num2 = num2, num1
	}
	res := ""
	carry := 0
	for idx := 0; idx < len(num1); idx += 1 {
		numBit1 := int(num1[len(num1)-idx-1] - '0')
		numBit2 := int(num2[len(num2)-idx-1] - '0')
		res = strconv.Itoa((numBit1+numBit2+carry)%10) + res
		carry = (numBit1 + numBit2 + carry) / 10
	}
	for idx := len(num2) - len(num1) - 1; idx >= 0; idx -= 1 {
		numBit2 := int(num2[idx] - '0')
		res = strconv.Itoa((numBit2+carry)%10) + res
		carry = (numBit2 + carry) / 10
	}
	if carry > 0 {
		res = strconv.Itoa(carry) + res
	}

	return res
}

// @lc code=end

