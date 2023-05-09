/*
 * @lc app=leetcode.cn id=521 lang=golang
 *
 * [521] 最长特殊序列 Ⅰ
 */

// @lc code=start
func findLUSlength(a string, b string) int {
	// solution: 若两串字符串不等，则较长的字符串一定不会是较短字符串的子序列，若两串相同则返回-1
	if a != b {
		if len(a) > len(b) {
			return len(a)
		} else {
			return len(b)
		}
	}

	return -1
}

// @lc code=end

