/*
 * @lc app=leetcode.cn id=551 lang=golang
 *
 * [551] 学生出勤记录 I
 */

// @lc code=start
func checkRecord(s string) bool {
	// solution: 模拟
	aCounter := 0
	for idx := 0; idx < len(s); idx += 1 {
		switch s[idx] {
		case 'A':
			aCounter += 1
			if aCounter == 2 {
				return false
			}
		case 'L':
			if idx+3 <= len(s) && s[idx:idx+3] == "LLL" {
				return false
			}
		case 'P':
			continue
		}
	}

	return true
}

// @lc code=end

