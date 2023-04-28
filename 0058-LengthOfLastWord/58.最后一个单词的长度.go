/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	maxLength := 0
	for idx := len(s) - 1; idx >= 0; idx -= 1 {
		if s[idx] == ' ' {
			if maxLength > 0 {
				break
			}
			continue
		}
		maxLength += 1
	}
	return maxLength
}

// @lc code=end

