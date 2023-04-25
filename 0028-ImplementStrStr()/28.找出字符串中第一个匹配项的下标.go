/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 找出字符串中第一个匹配项的下标
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	for idx := 0; idx <= len(haystack)-len(needle); idx += 1 {
		if haystack[idx:idx+len(needle)] == needle {
			return idx
		}
	}
	return -1
}

// @lc code=end

