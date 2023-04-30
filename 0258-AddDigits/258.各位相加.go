/*
 * @lc app=leetcode.cn id=258 lang=golang
 *
 * [258] 各位相加
 */

// @lc code=start
func addDigits(num int) int {
	if num < 10 {
		return num
	}
	for num = getNumSum(num); num >= 10; num = getNumSum(num) {

	}
	return num
}

func getNumSum(num int) int {
	sum := 0
	for ; num > 0; num /= 10 {
		sum += num % 10
	}
	return sum
}

// @lc code=end

