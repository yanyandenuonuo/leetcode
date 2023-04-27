/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	// 反转一半即可
	reverseNum := 0
	for x > reverseNum {
		reverseNum = reverseNum*10 + x%10
		if x/10 >= reverseNum {
			x /= 10
		}
	}

	return x == reverseNum
}

// @lc code=end

