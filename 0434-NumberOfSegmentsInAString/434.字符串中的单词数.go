/*
 * @lc app=leetcode.cn id=434 lang=golang
 *
 * [434] 字符串中的单词数
 */

// @lc code=start
func countSegments(s string) int {
	counter := 0
	for idx := 0; idx < len(s); idx += 1 {
		if s[idx] != ' ' && (idx == 0 || s[idx-1] == ' ') {
			counter += 1
		}
	}
	return counter
}

// @lc code=end

