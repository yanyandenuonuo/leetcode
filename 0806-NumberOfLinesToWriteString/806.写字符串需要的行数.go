/*
 * @lc app=leetcode.cn id=806 lang=golang
 *
 * [806] 写字符串需要的行数
 */

// @lc code=start
func numberOfLines(widths []int, s string) []int {
	lineCount, lineIdx := 100, 1

	for _, runeChar := range s {
		charCount := widths[runeChar-'a']

		if charCount > lineCount {
			lineIdx += 1
			lineCount = 100 - charCount
		} else {
			lineCount -= charCount
		}
	}

	return []int{lineIdx, 100 - lineCount}
}

// @lc code=end

