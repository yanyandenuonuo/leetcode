/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
func romanToInt(s string) int {
	romanToIntMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	sum := 0
	leftSum := 0
	for _, sBit := range s {
		intBit := romanToIntMap[sBit]
		if leftSum < intBit {
			sum = sum - leftSum + intBit - leftSum
			leftSum = intBit
		} else if leftSum == intBit {
			leftSum += intBit
			sum += intBit
		} else {
			leftSum = intBit
			sum += intBit
		}
	}
	return sum
}

// @lc code=end

