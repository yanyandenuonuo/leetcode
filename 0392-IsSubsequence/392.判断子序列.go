/*
 * @lc app=leetcode.cn id=392 lang=golang
 *
 * [392] 判断子序列
 */

// @lc code=start
func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	for idx := 0; idx < len(t); idx += 1 {
		if s[0:1] == t[idx:idx+1] {
			s = s[1:]
		}
		if len(s) == 0 {
			return true
		}
	}
	return false
}

// @lc code=end

