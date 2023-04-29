/*
 * @lc app=leetcode.cn id=171 lang=golang
 *
 * [171] Excel 表列序号
 */

// @lc code=start
func titleToNumber(columnTitle string) int {
	columnNum := 0
	for _, runeChar := range columnTitle {
		columnNum = columnNum*26 + int(runeChar-'A'+1)
	}
	return columnNum
}

// @lc code=end

